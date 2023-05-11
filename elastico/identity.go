/*
 * Copyright © 2018 Lynn <lynn9388@gmail.com>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package elastico

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"log"
	"math"
	"strconv"
	"time"

	"github.com/harmony-one/vdf/src/vdf_go"

	//"github.com/stretchr/testify/assert"
	"github.com/lynn9388/pox/pow"
)

const (
	random     = "lynn9388"
	difficulty = 10000
	shardNum   = 4
)

// NewIDProof returns a new proof for identity with PoW.
func NewIDProof(addr string, pk []byte) *IDProof {
	// var data bytes.Buffer
	// data.WriteString(random)
	// data.WriteString(addr)
	// data.Write(pk)
	//测试
	input := [32]byte{}
	rand.Read(input[:])
	vdf := vdf_go.New(difficulty, input)
	outputChannel := vdf.GetOutputChannel()
	start := time.Now()
	vdf.Execute()
	duration := time.Now().Sub(start)
	log.Printf("VDF computation finished, time spent %s", duration.String())
	output := <-outputChannel
	nonce := output[:]
	//log.Printf("output=%s", string(output[:]))
	log.Println("finshed nonce set!")
	return &IDProof{
		Addr:  addr,
		PK:    pk,
		Nonce: nonce,
		Vdf:   *vdf,
	}
}

// toByte converts the identity proof to slice of byte without nonce field.
func (p *IDProof) toByte() []byte {
	var data bytes.Buffer
	data.WriteString(random)
	data.WriteString(p.Addr)
	data.Write(p.PK)
	// err := binary.Write(&data, binary.LittleEndian, difficulty)
	// if err != nil {
	// 	log.Println("binary.Write failed", err)
	// }
	return data.Bytes()
}

// Verify verifies if the identity proof is correct.
func (p *IDProof) Verify(vdf vdf_go.VDF) bool {
	output := vdf.GetOutput()
	if vdf.Verify(output) {
		log.Println("verify true!")
		return true
	}
	log.Println("verify false!")
	return false
}

// GetCommitteeNo generates the committee number based on the identity proof.
func (p *IDProof) GetCommitteeNo() int {
	nonce := int64(binary.BigEndian.Uint64(p.Nonce))
	hash := pow.Hash(p.toByte(), nonce)
	bitNum := int(math.Ceil(math.Log2(shardNum)))
	//log.Printf("接下来是bitnum:%d   %d", len(hash), len(hash)-bitNum)
	lastBytes := hash[len(hash)-bitNum:]
	value, err := strconv.ParseUint(hex.EncodeToString(lastBytes[:]), 16, 8*bitNum)
	if err != nil {
		log.Printf("接下来是错误:%d", value)
		log.Fatal(err)
	}
	log.Println("committeeNo finshed!")
	return int(value) % shardNum
}

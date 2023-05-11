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
	// "math/rand"
	//"crypto/rand"
	"encoding/binary"
	"log"
	"testing"

	//"encoding/hex"
	//"github.com/harmony-one/vdf/src/vdf_go"
	"github.com/lynn9388/blockchain-sharding/crypto"
	"github.com/lynn9388/pox/pow"
)

func TestIDProof_Verify(t *testing.T) {
	sk, err := crypto.NewKey()
	if err != nil {
		t.Fatal(err)
	}
	proof:= NewIDProof("localhost:9388", sk.D.Bytes())
	//output := vdf.GetOutput()
	//log.Printf("output=%s", string(output[:]))
	if proof.Verify(proof.Vdf) != true {
		log.Println("this have error!!!!")
		t.Fail()
	}
	//proof.Addr = "localhost:9488"
	// if proof.Verify(*vdf) == true {
	// 	t.Fail()
	// 	log.Println("this have error!")
	// }
	log.Println("finshed TestIDProof_verify!")
}

func TestIDProof_GetCommitteeNo(t *testing.T) {
	sk, err := crypto.NewKey()
	if err != nil {
		t.Fatal(err)
	}
	proof:= NewIDProof("localhost:9388", sk.D.Bytes())
	nonce := int64(binary.BigEndian.Uint64(proof.Nonce))
	hash := pow.Hash(proof.toByte(), nonce)
	log.Println("hash finshed!")
	no := proof.GetCommitteeNo()
	log.Printf("no is %d:", no)
	t.Logf("committee number: %v -> %v", hash, no)
	log.Printf("committee number: %v -> %v", hash, no)
}

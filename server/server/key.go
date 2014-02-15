//
// Author Cloud@txthinking.com
//
package server

import(
    "time"
    "math/rand"
    "crypto/sha1"
    "encoding/hex"
    "hash"
    "io"
    "strconv"
)

func Key()(s string){
    var i0 int64 = time.Now().UnixNano()
    var i1 int64 = rand.New(rand.NewSource(i0)).Int63()
    var h hash.Hash = sha1.New()
    io.WriteString(h, strconv.FormatInt(i0, 10))
    io.WriteString(h, strconv.FormatInt(i1, 10))
    s = hex.EncodeToString(h.Sum(nil))
    return
}


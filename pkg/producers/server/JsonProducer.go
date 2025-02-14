// Copyright © 2024 JR team
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package server

import (
	"net/http"

	"github.com/jrnd-io/jr/pkg/tpl"
	"github.com/rs/zerolog/log"
)

type JsonProducer struct {
	OutputTpl *tpl.Tpl
}

func (c *JsonProducer) Close() error {
	// no need to close
	return nil
}

func (c *JsonProducer) Produce(key []byte, value []byte, o interface{}) {

	if o != nil {
		respWriter := o.(http.ResponseWriter)
		if string(key) != "null" {
			_, err := (respWriter).Write(key)
			if err != nil {
				log.Error().Err(err).Msg("Error writing key")
			}
			_, err = (respWriter).Write([]byte(","))
			if err != nil {
				log.Error().Err(err).Msg("Error writing comma")
			}
		}
		_, err := (respWriter).Write(value)
		if err != nil {
			log.Error().Err(err).Msg("Error writing value")
		}
	} else {
		log.Warn().Interface("o", o).Msg("Server producer must produce to a http.ResponseWriter")
	}

}

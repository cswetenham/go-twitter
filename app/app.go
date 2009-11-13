//
// Copyright 2009 Bill Casarin <billcasarin@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
package main

import "twitter"
import "fmt"

func main() {
  api := twitter.NewApi();

  api.Authenticate("jb55", "");
  statusChan := api.GetStatusAsync(5641609144);
  statuses := api.GetUserTimeline();
  for i, status := range statuses {
    fmt.Printf("#%d: %s\n", i, status.GetText());
  }
  fmt.Printf("GetStatusAsync: %s\n", (<-statusChan).GetText());
  //api.PostUpdate("Testing my Go twitter library", 0);
}

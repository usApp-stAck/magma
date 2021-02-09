/**
 * Copyright 2020 The Magma Authors.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

#include "EventBaseWrapper.h"

namespace magma {

EventBaseWrapper::EventBaseWrapper(folly::EventBase* evb) : evb_(evb) {}

void EventBaseWrapper::loopForever() {
  evb_->loopForever();
};

void EventBaseWrapper::terminateLoopSoon() {
  evb_->terminateLoopSoon();
};

void EventBaseWrapper::runAfterDelay(
    folly::Function<void()> func, int32_t delayMs) {
  evb_->runAfterDelay(std::move(func), delayMs);
};

void EventBaseWrapper::runInEventBaseThread(folly::Cob&& cob) {
  evb_->runInEventBaseThread(std::move(cob));
}
};  // namespace magma
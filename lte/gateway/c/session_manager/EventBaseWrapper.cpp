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
#include "magma_logging.h"
#include "Utilities.h"
#include "date.h"

namespace magma {
using namespace date;

EventBaseWrapper::EventBaseWrapper(folly::EventBase* evb) : evb_(evb) {}

void EventBaseWrapper::loopForever() {
  evb_->loopForever();
};

void EventBaseWrapper::terminateLoopSoon() {
  evb_->terminateLoopSoon();
};

void EventBaseWrapper::runAfterDelay(
    folly::Function<void()> func, int32_t delayMs) {
  auto now = std::chrono::system_clock::now();
  log_push(now);
  evb_->runAfterDelay(
      [this, func = std::move(func), now]() mutable {
        func();
        log_pop(now);
      },
      delayMs);
};

void EventBaseWrapper::runInEventBaseThread(folly::Cob&& cob) {
  auto now = std::chrono::system_clock::now();
  log_push(now);
  evb_->runInEventBaseThread([this, func = std::move(cob), now]() mutable {
    func();
    log_pop(now);
  });
}

void EventBaseWrapper::log_push(std::chrono::system_clock::time_point now) {
  event_count++;
  MLOG(MINFO) << "[" << event_count << "] ==> Inserting into EventQueue at "
              << now;
};

void EventBaseWrapper::log_pop(std::chrono::system_clock::time_point then) {
  auto now = std::chrono::system_clock::now();
  event_count--;
  MLOG(MINFO) << "[" << event_count << "] "
              << "<== Popping from EventQueue (inserted: " << then
              << ", duration: " << (now - then);
};
};  // namespace magma
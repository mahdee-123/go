text akare dewa

3:34 PM
Section 1 — Basics

# divide(a, b int) (int, error) — b==0 হলে error, properly check করে call করো।
# parseAge(s string) (int, error) — strconv.Atoi ব্যবহার করে, fmt.Errorf দিয়ে %w wrap করো।
readConfig() — file read→parse→validate, early return pattern দিয়ে (nested if ছাড়া)।
_ দিয়ে error ignore করার example লেখো, explain করো কেন risky।
Section 2 — Custom errors

NotFoundError struct (Resource, ID) + Error() method। findUser(id) এ ব্যবহার করো।
Sentinel error ErrNotFound বানাও, wrap করো, errors.Is() দিয়ে check করো।
উপরের NotFoundError wrap করে errors.As() দিয়ে data বের করো।
৩ layer chain (repo→service→handler), প্রতিটায় wrap, errors.Unwrap() দিয়ে layer খোলো।
দুটো sentinel error (ErrInvalidAge, ErrEmptyName) দিয়ে validation function।
Section 3 — Panic, recover, defer

একটা panic example (nil map write/index out of range) — explain panic vs error।
safeDivide — panic হলে recover করে error এ convert করো।
file open + defer Close(), named return এ close error capture করো।
goroutine এ panic — main এর recover কাজ করে না দেখাও, তারপর fix করো।
Section 4 — Context & timeouts

context.WithTimeout দিয়ে slow operation, ctx.Err() চেক করো।
context.DeadlineExceeded vs context.Canceled — errors.Is দিয়ে আলাদা handle।
৩ goroutine + error channel — সব error collect করো।
errgroup দিয়ে ৩ goroutine, একটা fail করলে সব cancel।
Section 5 — Real backend scenarios

HTTP handler — error type অনুযায়ী status code (404/400/500) ও JSON response।
callExternalAPI() — exponential backoff সহ 3 retry।
logErrors(err) — Unwrap loop দিয়ে full chain log, user কে শুধু top message।
;; Define and validate complex expressions
(set-logic QF_AUFLIA)

;; Array and arithmetic operations
(declare-fun arr () (Array Int Int))
(declare-fun x () Int)
(declare-fun y () String)

;; Combine multiple theories in assertions
(assert (> x 0))
(assert (= (select arr x) (* 2 x)))
(assert (str.contains y "hello"))

(check-sat)
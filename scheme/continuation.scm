(define (sum x y) (+ x y))
(sum 2 3)
(define (square x) (* x x))
(define (tri x y) (sum (square x) (square y)))
(tri 3 4)

(define l
  (for-each
   (lambda (x) (* x x))
   '(1 2 3 4 5)))
(map
 (lambda (x) (* x x))
 '(1 2 3 4 5))


;;; example 1: non-local exit
(define (test element cc)
  (if (zero? element)
      (cc 'found-zero)
      (void)))
(define (search-zero test lst)
  (call/cc
   (lambda (return)
     (for-each (lambda (element)
		 (test element return)
		 (printf "~a~%" element))
	       lst)
     #f)))
(search-zero test '(1 2 3 0 1 2 3))

(define cc #f)
(+ 1 (call/cc
      (lambda (k)
	(set! cc k)
	1)))
(cc 2)
;;; end


;;; example 2
(define (hefty-computation do-other-stuff) 
    (let loop ((n 5)) 
      (display "Hefty computation: ") 
      (display n) 
      (newline) 
      (set! do-other-stuff (call/cc do-other-stuff)) 
      ;; (call/cc do-other-stuff)
      (display "Hefty computation (b)")  
      (newline) 
      ;; (call/cc do-other-stuff)
      (set! do-other-stuff (call/cc do-other-stuff)) 
      (display "Hefty computation (c)") 
      (newline) 
      ;; (call/cc do-other-stuff)
      (set! do-other-stuff (call/cc do-other-stuff)) 
      (if (> n 0) 
          (loop (- n 1))
	  #t))) 
(define (superfluous-computation do-other-stuff) 
    (let loop () 
      (for-each (lambda (graphic) 
                  (display graphic) 
                  (newline) 
                  (set! do-other-stuff (call/cc do-other-stuff))) 
                '("Straight up." "Quarter after." "Half past."  "Quarter til.")) 
      (loop))) 
(hefty-computation superfluous-computation)
;;; end of example 2 



;;; example 3: generator
(define (generate-one-element-at-a-time lst)
  (define (control-state return)
    (for-each
     (lambda (element)
       (call/cc
	(lambda (resume-here)
	  (set! control-state resume-here)
	  (return element))))
     lst)
    (return 'you-fell-off-the-end))
  (define (generator)
    (call/cc control-state))
  generator)

(define generate-digit
  (generate-one-element-at-a-time '(0 1 2)))

(generate-digit)
;;; end of example 3: generator



;;; define a function returning the current continuation
(define (get-cc) (call/cc (lambda (c) c)))
(define x (get-cc))
(x 10)
;;; end


;;; coroutine: 使用延续模仿多任务
(define lwp-list '())
(define lwp
  (lambda (thunk)
    (set! lwp-list (append lwp-list (list thunk)))))

(define start
  (lambda ()
    (let ([p (car lwp-list)])
      (set! lwp-list (cdr lwp-list))
      (p))))

(define pause
  (lambda ()
    (call/cc
      (lambda (k)
        (lwp (lambda () (k #f)))
        (start)))))

(lwp (lambda () (let f () (pause) (display "h") (f))))
(lwp (lambda () (let f () (pause) (display "e") (f))))
(lwp (lambda () (let f () (pause) (display "y") (f))))
(lwp (lambda () (let f () (pause) (display "!") (f))))
(lwp (lambda () (let f () (pause) (newline) (f))))
;;; end



;;; the most difficult example to understand 
(let * ((yin
	 ((lambda (cc) (display #\@) cc) (call/cc (lambda (c) c))))
	(yang
	 ((lambda (cc) (display #\*) cc) (call/cc (lambda (c) c)))))(yin yang))
;;; end

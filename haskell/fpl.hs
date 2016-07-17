{-module Main where-}

sumTo20 :: [Int] -> Int
sumTo20 nums = go 0 nums
  where go :: Int -> [Int] -> Int
        go acc [] = acc
        go acc (x:xs)
         | acc >= 20 = acc
         | otherwise = go (acc + x) xs

fibs :: [Integer]
fibs = 0 : 1 : [a + b | (a, b) <- zip fibs (tail fibs)]

mysumAll :: [Int] -> Int
mysumAll [] = 0
mysumAll (x:xs) = x + mysumAll xs

sum' :: (Num a) => [a] -> a
sum' [] = 0
sum' (x:xs) = x + sum' xs

qsort :: (Ord a) => [a] -> [a]
qsort [] = []
qsort (x:xs) = 
    let smallerOrEqual = [y | y <- xs, y <= x]
        bigger = [y | y <- xs, y > x]
     in qsort smallerOrEqual ++ [x] ++ qsort bigger

doubleSmallNumber x = if x<100
                         then x+x
                         else x

len' :: [a] -> Int
len' x = sum' [1 | a <- x]

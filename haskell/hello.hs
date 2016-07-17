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

mysum :: (Num a) => [a] -> a
mysum [] = 0
mysum (x:xs) = x + mysum xs

quickSort :: (Ord a) => [a] -> [a]
quickSort [] = []
quickSort (x:xs) = (quickSort [y | y <- xs, y <= x] ++ [x] ++ quickSort [y | y <- xs, y > x])

doubleSmallNumber x = if x<100
                         then x+x
                         else x

tripleSmallNum x = if x<100
	then x+x+x
	else x

length x = sum [1 | _ <- x]

{-main :: IO ()-}
{-main = do-}
  {-print "hello world"-}
  {-sumTo20 [1,2,3,4,5,6,7,8,9]-}

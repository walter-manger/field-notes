
doubleMe x = x + x

doubleUs x y = doubleMe x + doubleMe y

doubleSmallNumber x = if x > 100
  then x
  else x*2

removeUppercase :: String -> String
removeUppercase st = [c | c <- st, c `elem` ['a'..'z']]

addThree :: Int -> Int -> Int -> Int
addThree x y z = x + y + z

-- Syntax In Functions

lucky :: (Integral a) => a -> String
lucky 7 = "LUCKY NUMBER SEVEN!"
lucky x = "Sorry, you're out of luck, pal!"

factorial :: (Integral a) => a -> a
factorial 0 = 1
factorial n = n * factorial (n - 1)

addVectors :: (Num a) => (a, a) -> (a, a) -> (a, a)
-- addVectors a b = (fst a + fst b, snd a + snd b)
-- pattern matching
addVectors (x1, y1) (x2, y2) = (x1 + x2, y1 + y2)


first :: (a, b, c) -> a
first (x, _, _) = x

second :: (a, b, c) -> b
second (_, y, _) = y

third :: (a, b, c) -> c
third (_, _, z) = z

head' :: [a] -> a
head' [] = error "Can't call head on an empty list, dummy!"
head' (x:_) = x
-- Need parens if we're binding more than one variable. x and _ make 2


tell :: (Show a) => [a] -> String
tell [] = "The list is empty"
-- (With sugar) tell [x] = "The list has one element: " ++ show x
tell (x:[]) = "The list has one element: " ++ show x
tell (x:y:[]) = "The list has two elements: " ++ show x ++ " and " ++ show y
tell (x:y:_) = "The list is long. The first two elements: " ++ show x ++ " and " ++ show y


length' :: (Num b) => [a] -> b
length' [] = 0
length' (_:xs) = 1 + length' xs


someSum :: (Num a) => [a] -> a
someSum [] = 0
someSum (x:xs) = x + someSum xs


-- bmiTell :: (RealFloat a) => a -> String
-- bmiTell bmi
--   | bmi <= 18.5 = "You're underweight, you emo, you!"
--   | bmi <= 25.0 = "You're supposedly normal. Pffft, I bet you're ugly!"
--   | bmi <= 30.0 = "You're fat. Lose some weight, fatty!"
--   | otherwise = "You're a whale, congratulations!"


-- bmiTell :: (RealFloat a) => a -> a -> String
-- bmiTell weight height
--   | weight / height ^ 2 <= 18.5 = "You're underweight, you emo, you!"
--   | weight / height ^ 2 <= 25.0 = "You're supposedly normal. Pffft, I bet you're ugly!"
--   | weight / height ^ 2 <= 30.0 = "You're fat. Lose some weight, fatty!"
--   | otherwise = "You're a whale, congratulations!"


max' :: (Ord a) => a -> a -> a
max' a b
  | a > b = a
  | otherwise = b


-- Backticks for readablility
myCompare :: (Ord a) => a -> a -> Ordering
a `myCompare` b
  | a > b = GT
  | a == b = EQ
  | otherwise = LT


-- where for readability
bmiTell :: (RealFloat a) => a -> a -> String
bmiTell weight height
  | bmi <= skinny = "You're underweight, you emo, you!"
  | bmi <= normal = "You're supposedly normal. Pffft, I bet you're ugly!"
  | bmi <= fat = "You're fat. Lose some weight, fatty!"
  | otherwise = "You're a whale, congratulations!"
  where bmi = weight / height ^2
        (skinny, normal, fat) = (18.5, 25.0, 30.0) -- Awesome pattern matching
        -- skinny = 18.5
        -- normal = 25.0
        -- fat = 30.0


initials :: String -> String -> String
initials first last = [f] ++ ". " ++ [l] ++ "."
  where (f:_) = first
        (l:_) = last


cylinder :: (RealFloat a) => a -> a -> a
cylinder r h =
  let sideArea = 2 * pi * r * h
      topArea = pi * r * 2
  in sideArea + 2 * topArea

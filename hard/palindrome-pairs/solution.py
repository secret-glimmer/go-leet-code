class Solution:
    def palindromePairs(self, words: List[str]) -> List[List[int]]:
        word_dict = { word: index for index, word in enumerate(words) }
        result = []

        for index, word in enumerate(words):
            for j in range(len(word) + 1):
                left = word[:j]
                reversed_left = left[::-1]
                right = word[j:]
                reversed_right = right[::-1]

                if reversed_left in word_dict \
                    and word_dict[reversed_left] != index \
                    and right == reversed_right:
                    result.append([index, word_dict[reversed_left]])
                
                if j > 0 \
                    and reversed_right in word_dict \
                    and word_dict[reversed_right] != index \
                    and left == reversed_left:
                    result.append([word_dict[reversed_right], index])
            
        
        return result

        
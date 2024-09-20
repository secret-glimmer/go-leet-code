class Solution:
    def shortestPalindrome(self, s):
        # Get the length of the longest palindromic prefix
        count = self.kmp(s[::-1], s)
        # Append the non-palindromic suffix in reverse to the front
        return s[count:][::-1] + s

    def kmp(self, txt, patt):
        # Create the combined string for KMP processing
        new_string = patt + '#' + txt
        pi = [0] * len(new_string)
        i = 1
        k = 0
        
        # Build the prefix table
        while i < len(new_string):
            if new_string[i] == new_string[k]:
                k += 1
                pi[i] = k
                i += 1
            else:
                if k > 0:
                    k = pi[k - 1]
                else:
                    pi[i] = 0
                    i += 1
        
        # Return the length of the longest palindromic prefix
        return pi[-1]

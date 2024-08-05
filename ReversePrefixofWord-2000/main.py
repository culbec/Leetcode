class Solution(object):
    def reversePrefix(self, word, ch):
        """
        :type word: str
        :type ch: str
        :rtype: str
        """
        i = word.find(ch)
        if i == -1:
            return word

        return word[:i+1][::-1] + word[i+1:]

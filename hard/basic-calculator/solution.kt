class Solution {
    fun calculate(s: String): Int {
        var result = 0 // This will store the final result of the expression
        var num = 0 // This will store the current number being processed
        var sign = 1 // This keeps track of the current sign (+1 for positive, -1 for negative)
        val bracketSigns = LinkedList<Int>() // This stack will store the sign before encountering a '('
        var bracketSign = 1 // This keeps track of the cumulative sign for the current expression inside parentheses
        // Function to handle digits and update the result when a complete number is formed
        fun numHandler(ch: Char) {
            when {
                ch in '0'..'9' -> { 
                    // If the character is a digit, update the current number
                    num = num * 10 + (ch - '0')
                }
                num > 0 -> {
                    // If the current character is not a digit, and a number was previously being formed,
                    // add it to the result, adjusting for the sign and bracketSign, then reset `num`
                    result += sign * num * bracketSign
                    num = 0
                }
            }
        }
        // Function to handle signs and parentheses
        fun signHandler(ch: Char) {
            when (ch) {
                '-' -> sign = -1 // If the character is '-', set the sign to negative
                '+' -> sign = 1 // If the character is '+', set the sign to positive
                ')' -> bracketSign = bracketSigns.pop() // If a ')' is encountered, restore the previous bracket sign
                '(' -> {
                    // If a '(' is encountered, push the current bracket sign onto the stack
                    // and update the bracketSign by multiplying it with the current sign
                    bracketSigns.push(bracketSign)
                    bracketSign *= sign
                    sign = 1 // Reset the sign to positive after opening a new parenthesis
                }
            }
        }
        // Loop through each character in the input string
        for (ch in s) {
            if (ch == ' ') {
                continue // Skip whitespace characters
            }
            numHandler(ch) // Process the character as a number if applicable
            signHandler(ch) // Process the character as a sign or parenthesis if applicable
        }
        // After the loop, add the last number to the result
        return result + sign * num * bracketSign
    }
}
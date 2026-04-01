// Chibueze Maxwell

package processor

import (
 "strings"
)

// FixQuote removes the spaces immediately after the opening '
// and before the closing ' while keeping everything inside intact.
func FixQuote(s string) string {
 var result strings.Builder
 inQuote := false
 quoteStart := -1

 for i := 0; i < len(s); i++ {
  if s[i] == '\'' {
   // First quote found
   if !inQuote {
    inQuote = true
    quoteStart = len(result.String())
    result.WriteByte('\'')
   } else {
    // Closing quote found
    inQuote = false

    // Trim space directly before closing quote
    temp := result.String()
    if len(temp) > 0 && temp[len(temp)-1] == ' ' {
     // remove the last space
     result.Reset()
     result.WriteString(temp[:len(temp)-1])
    }

    result.WriteByte('\'')
   }
  } else {
   // Characters inside a quote
   if inQuote {
    // Skip leading spaces right after first quote
    if len(result.String()) == quoteStart+1 && s[i] == ' ' {
     continue
    }
    result.WriteByte(s[i])
   } else {
    result.WriteByte(s[i])
   }
  }
 }

 return result.String()
}
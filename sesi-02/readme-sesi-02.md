# SESI II

# Summary

- Condition:
  - if-else
  - switch
    - fallthrough >> melanjutkan pengecekan kepada case selanjutnya walaupun suatu case telah terpenuhi kondisinya
  - temporary variable >> hanya bisa diakses dan digunakan pada scope block dari suatu kondisional.
- Looping:
  - Bahasa Go hanya memiliki satu looping yaitu looping dengan menggunakan keyword for
  - label >> memberikan tanda pada proses looping
- Array:
  - Array pada bahasa Go memiliki sifat fixed-length
- Slice:
  - tidak memiliki sifat fixed-length
  - reference type
  - make(type, length) >> create slice
  - append() >> add item
  - copy(to, from)
  - slicing \[fromIndex : toIndex\]
  - cap() >> capacity
  - len() >> length
- Map:
  - key:value pairs
  - map[keyType]valueType
  - delete(var, key) >> delete value
  - detecting value with variable exist
- Aliase:
  - nama alternative dari tipe data yang sudah ada
  - byte alias dari uint8
  - rune alias dari uint32
- Strings in depth:
  - Ketika kita melakukan indexing terhadap suatu string, maka kita akan mendapat nilai representasi dari byte nya.
  - len() >> check number of byte
  - utf8.RuneCountInString() >> length of string
  -

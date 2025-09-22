; ModuleID = 'add_test.uni'
source_filename = "add_test.uni"

define i32 @add(i32 %a, i32 %b) {
entry:
  %add = add i32 %a, %b
  ret i32 %add
}

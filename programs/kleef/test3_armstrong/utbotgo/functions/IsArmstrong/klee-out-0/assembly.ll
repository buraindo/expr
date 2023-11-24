; ModuleID = 'utbotgo/functions/IsArmstrong/llvm_bitcode.ll'
source_filename = "gomodule"
target datalayout = "e-m:e-p270:32:32-p271:32:32-p272:64:64-i64:64-f80:128-n8:16:32:64-S128"
target triple = "x86_64-unknown-linux-gnu"

module asm "\09.section \22.go_export\22,\22e\22,@progbits"
module asm "\09.ascii \22v3;\\n\22"
module asm "\09.ascii \22package \22"
module asm "\09.ascii \22armstrong\22"
module asm "\09.ascii \22\\n\22"
module asm "\09.ascii \22pkgpath \22"
module asm "\09.ascii \22test3\22"
module asm "\09.ascii \22\\n\22"
module asm "\09.ascii \22init\22"
module asm "\09.ascii \22 \22"
module asm "\09.ascii \22armstrong\22"
module asm "\09.ascii \22 \22"
module asm "\09.ascii \22~test3\22"
module asm "\09.ascii \22\\n\22"
module asm "\09.ascii \22types 1 1\22"
module asm "\09.ascii \22\\n\22"
module asm "\09.ascii \22\22"
module asm "\09.ascii \22func \22"
module asm "\09.ascii \22IsArmstrong\22"
module asm "\09.ascii \22 (\22"
module asm "\09.ascii \22n\22"
module asm "\09.ascii \22 \22"
module asm "\09.ascii \22<type -11>\22"
module asm "\09.ascii \22)\22"
module asm "\09.ascii \22 \22"
module asm "\09.ascii \22<type -15>\22"
module asm "\09.ascii \22\\n\22"
module asm "\09.ascii \22checksum FEB85A1AED8AA09E28463AB73E7CCEB909193951\\n\22"
module asm "\09.text"

%functionDescriptor.0 = type { i64 }

@test3.IsArmstrong..f = constant %functionDescriptor.0 { i64 ptrtoint (i8 (i8*, i64)* @test3.IsArmstrong to i64) }
@test3..types = constant { i64, [1 x i8*] } zeroinitializer
@klee.zero = internal constant i64 0
@klee.n.str = internal constant [2 x i8] c"n\00"
@klee_modify..klee.n.str = internal constant i8 2

declare void @0(i8*, ...)

define internal i64 @test3.pow(i8* nest %nest.0, i64 %a, i64 %p) #0 !dbg !5 {
entry:
  %a.addr = alloca i64, align 8
  %p.addr = alloca i64, align 8
  %"$ret0" = alloca i64, align 8
  %r = alloca i64, align 8
  %i = alloca i64, align 8
  store i64 %a, i64* %a.addr, align 8
  call void @llvm.dbg.declare(metadata i64* %a.addr, metadata !11, metadata !DIExpression()), !dbg !12
  store i64 %p, i64* %p.addr, align 8
  call void @llvm.dbg.declare(metadata i64* %p.addr, metadata !13, metadata !DIExpression()), !dbg !14
  %0 = bitcast i64* %"$ret0" to i8*
  store i64 0, i64* %"$ret0", align 8
  call void @llvm.dbg.declare(metadata i64* %"$ret0", metadata !15, metadata !DIExpression()), !dbg !16
  %1 = bitcast i64* %r to i8*
  %a.ld.0 = load i64, i64* %a.addr, align 8, !dbg !17
  store i64 %a.ld.0, i64* %r, align 8
  call void @llvm.dbg.declare(metadata i64* %r, metadata !19, metadata !DIExpression()), !dbg !20
  %2 = bitcast i64* %i to i8*
  store i64 0, i64* %i, align 8
  call void @llvm.dbg.declare(metadata i64* %i, metadata !21, metadata !DIExpression()), !dbg !23
  br label %label.0

label.0:                                          ; preds = %then.0, %entry
  %i.ld.1 = load i64, i64* %i, align 8, !dbg !24
  %p.ld.0 = load i64, i64* %p.addr, align 8, !dbg !25
  %sub.0 = sub i64 %p.ld.0, 1, !dbg !26
  %icmp.0 = icmp slt i64 %i.ld.1, %sub.0, !dbg !27
  %zext.0 = zext i1 %icmp.0 to i8, !dbg !27
  %trunc.0 = trunc i8 %zext.0 to i1, !dbg !27
  br i1 %trunc.0, label %then.0, label %else.0

then.0:                                           ; preds = %label.0
  %r.ld.0 = load i64, i64* %r, align 8, !dbg !28
  %a.ld.1 = load i64, i64* %a.addr, align 8, !dbg !29
  %mul.0 = mul i64 %r.ld.0, %a.ld.1, !dbg !30
  store i64 %mul.0, i64* %r, align 8, !dbg !30
  %i.ld.0 = load i64, i64* %i, align 8, !dbg !31
  %add.0 = add i64 %i.ld.0, 1, !dbg !31
  store i64 %add.0, i64* %i, align 8, !dbg !31
  br label %label.0

else.0:                                           ; preds = %label.0
  %3 = bitcast i64* %i to i8*
  %r.ld.1 = load i64, i64* %r, align 8, !dbg !32
  store i64 %r.ld.1, i64* %"$ret0", align 8, !dbg !33
  %"$ret0.ld.0" = load i64, i64* %"$ret0", align 8, !dbg !33
  %4 = bitcast i64* %r to i8*, !dbg !33
  %5 = bitcast i64* %"$ret0" to i8*, !dbg !33
  ret i64 %"$ret0.ld.0", !dbg !33
}

; Function Attrs: argmemonly nofree nosync nounwind willreturn
declare void @llvm.lifetime.start.p0i8(i64 immarg, i8* nocapture) #1

; Function Attrs: nofree nosync nounwind readnone speculatable willreturn
declare void @llvm.dbg.declare(metadata, metadata, metadata) #2

; Function Attrs: argmemonly nofree nosync nounwind willreturn
declare void @llvm.lifetime.end.p0i8(i64 immarg, i8* nocapture) #1

define i8 @test3.IsArmstrong(i8* nest %nest.1, i64 %n) #0 !dbg !34 {
entry:
  %n.addr = alloca i64, align 8
  %"$ret1" = alloca i8, align 1
  %rightMost = alloca i64, align 8
  %sum = alloca i64, align 8
  %tempNum = alloca i64, align 8
  %length = alloca i64, align 8
  %tmp = alloca i64, align 8
  %tmpv.0 = alloca i64, align 8
  %tmpv.1 = alloca i64, align 8
  %tmpv.2 = alloca i64, align 8
  %tmpv.3 = alloca i64, align 8
  %tmpv.4 = alloca i64, align 8
  %tmpv.5 = alloca i64, align 8
  %tmpv.6 = alloca i64, align 8
  store i64 %n, i64* %n.addr, align 8
  call void @llvm.dbg.declare(metadata i64* %n.addr, metadata !38, metadata !DIExpression()), !dbg !39
  store i8 0, i8* %"$ret1", align 1
  call void @llvm.dbg.declare(metadata i8* %"$ret1", metadata !40, metadata !DIExpression()), !dbg !41
  %0 = bitcast i64* %rightMost to i8*
  %1 = bitcast i64* %sum to i8*
  %2 = bitcast i64* %tempNum to i8*
  %3 = bitcast i64* %length to i8*
  %4 = bitcast i64* %tmp to i8*
  store i64 0, i64* %rightMost, align 8
  call void @llvm.dbg.declare(metadata i64* %rightMost, metadata !42, metadata !DIExpression()), !dbg !44
  store i64 0, i64* %sum, align 8
  call void @llvm.dbg.declare(metadata i64* %sum, metadata !45, metadata !DIExpression()), !dbg !46
  %n.ld.0 = load i64, i64* %n.addr, align 8, !dbg !47
  store i64 %n.ld.0, i64* %tempNum, align 8
  call void @llvm.dbg.declare(metadata i64* %tempNum, metadata !48, metadata !DIExpression()), !dbg !49
  store i64 1, i64* %length, align 8
  call void @llvm.dbg.declare(metadata i64* %length, metadata !50, metadata !DIExpression()), !dbg !51
  %n.ld.1 = load i64, i64* %n.addr, align 8, !dbg !52
  store i64 %n.ld.1, i64* %tmp, align 8
  call void @llvm.dbg.declare(metadata i64* %tmp, metadata !53, metadata !DIExpression()), !dbg !54
  br label %label.0

label.0:                                          ; preds = %then.3, %entry
  %tmp.ld.2 = load i64, i64* %tmp, align 8, !dbg !55
  %icmp.3 = icmp sge i64 %tmp.ld.2, 10, !dbg !56
  %zext.3 = zext i1 %icmp.3 to i8, !dbg !56
  %trunc.3 = trunc i8 %zext.3 to i1, !dbg !56
  br i1 %trunc.3, label %then.3, label %label.2

then.3:                                           ; preds = %label.0
  %tmp.ld.1 = load i64, i64* %tmp, align 8, !dbg !57
  %div.0 = sdiv i64 %tmp.ld.1, 10, !dbg !58
  store i64 %div.0, i64* %tmpv.0, align 8
  %tmpv.0.ld.0 = load i64, i64* %tmpv.0, align 8, !dbg !58
  store i64 %tmpv.0.ld.0, i64* %tmpv.1, align 8, !dbg !58
  %tmpv.1.ld.0 = load i64, i64* %tmpv.1, align 8, !dbg !58
  store i64 %tmpv.1.ld.0, i64* %tmp, align 8, !dbg !58
  %length.ld.0 = load i64, i64* %length, align 8, !dbg !59
  %add.1 = add i64 %length.ld.0, 1, !dbg !59
  store i64 %add.1, i64* %length, align 8, !dbg !59
  br label %label.0

label.2:                                          ; preds = %label.0, %then.8
  %tempNum.ld.3 = load i64, i64* %tempNum, align 8, !dbg !60
  %icmp.8 = icmp sgt i64 %tempNum.ld.3, 0, !dbg !61
  %zext.8 = zext i1 %icmp.8 to i8, !dbg !61
  %trunc.8 = trunc i8 %zext.8 to i1, !dbg !61
  br i1 %trunc.8, label %then.8, label %else.8

then.8:                                           ; preds = %label.2
  %tempNum.ld.0 = load i64, i64* %tempNum, align 8, !dbg !62
  %mod.0 = srem i64 %tempNum.ld.0, 10, !dbg !63
  store i64 %mod.0, i64* %tmpv.2, align 8
  %tmpv.2.ld.0 = load i64, i64* %tmpv.2, align 8, !dbg !63
  store i64 %tmpv.2.ld.0, i64* %tmpv.3, align 8, !dbg !63
  %tmpv.3.ld.0 = load i64, i64* %tmpv.3, align 8, !dbg !63
  store i64 %tmpv.3.ld.0, i64* %rightMost, align 8, !dbg !64
  %rightMost.ld.0 = load i64, i64* %rightMost, align 8, !dbg !65
  %length.ld.1 = load i64, i64* %length, align 8, !dbg !66
  %call.0 = call i64 @test3.pow(i8* nest undef, i64 %rightMost.ld.0, i64 %length.ld.1), !dbg !67
  store i64 %call.0, i64* %tmpv.4, align 8
  %sum.ld.0 = load i64, i64* %sum, align 8, !dbg !68
  %tmpv.4.ld.0 = load i64, i64* %tmpv.4, align 8, !dbg !67
  %add.2 = add i64 %sum.ld.0, %tmpv.4.ld.0, !dbg !69
  store i64 %add.2, i64* %sum, align 8, !dbg !69
  %tempNum.ld.2 = load i64, i64* %tempNum, align 8, !dbg !70
  %div.1 = sdiv i64 %tempNum.ld.2, 10, !dbg !71
  store i64 %div.1, i64* %tmpv.5, align 8
  %tmpv.5.ld.0 = load i64, i64* %tmpv.5, align 8, !dbg !71
  store i64 %tmpv.5.ld.0, i64* %tmpv.6, align 8, !dbg !71
  %tmpv.6.ld.0 = load i64, i64* %tmpv.6, align 8, !dbg !71
  store i64 %tmpv.6.ld.0, i64* %tempNum, align 8, !dbg !71
  br label %label.2

else.8:                                           ; preds = %label.2
  %n.ld.2 = load i64, i64* %n.addr, align 8, !dbg !72
  %sum.ld.1 = load i64, i64* %sum, align 8, !dbg !73
  %icmp.9 = icmp eq i64 %n.ld.2, %sum.ld.1, !dbg !74
  %zext.9 = zext i1 %icmp.9 to i8, !dbg !74
  store i8 %zext.9, i8* %"$ret1", align 1, !dbg !75
  %"$ret1.ld.0" = load i8, i8* %"$ret1", align 1, !dbg !75
  %5 = bitcast i64* %rightMost to i8*, !dbg !75
  %6 = bitcast i64* %sum to i8*, !dbg !75
  %7 = bitcast i64* %tempNum to i8*, !dbg !75
  %8 = bitcast i64* %length to i8*, !dbg !75
  %9 = bitcast i64* %tmp to i8*, !dbg !75
  ret i8 %"$ret1.ld.0", !dbg !75
}

; Function Attrs: noreturn
declare void @runtime.panicdivide(i8*) #3

declare void @klee_make_symbolic(i8*, i64, i8*)

define internal void @klee.main() {
entry:
  %klee.n.addr = alloca i64, align 8
  %klee.n.bc = bitcast i64* %klee.n.addr to i8*
  call void @klee_make_symbolic(i8* %klee.n.bc, i64 8, i8* getelementptr inbounds ([2 x i8], [2 x i8]* @klee.n.str, i64 0, i64 0))
  %klee.n.ld = load i64, i64* %klee.n.addr, align 8
  %klee.call = call i8 @test3.IsArmstrong(i8* undef, i64 %klee.n.ld)
  ret void
}

attributes #0 = { "disable-tail-calls"="true" "frame-pointer"="none" "null-pointer-is-valid"="true" "split-stack" "target-cpu"="x86-64" "target-features"="+cx8,+fxsr,+mmx,+sse,+sse2,+x87" }
attributes #1 = { argmemonly nofree nosync nounwind willreturn }
attributes #2 = { nofree nosync nounwind readnone speculatable willreturn }
attributes #3 = { noreturn "disable-tail-calls"="true" "frame-pointer"="none" "null-pointer-is-valid"="true" "split-stack" "target-cpu"="x86-64" "target-features"="+cx8,+fxsr,+mmx,+sse,+sse2,+x87" }

!llvm.module.flags = !{!0, !1, !2}
!llvm.dbg.cu = !{!3}

!0 = !{i32 7, !"PIC Level", i32 0}
!1 = !{i32 2, !"Debug Info Version", i32 3}
!2 = !{i32 2, !"Dwarf Version", i32 4}
!3 = distinct !DICompileUnit(language: DW_LANG_Go, file: !4, producer: "llvm-goc", isOptimized: true, runtimeVersion: 0, emissionKind: FullDebug)
!4 = !DIFile(filename: "./isarmstrong.go", directory: "/workspace")
!5 = distinct !DISubprogram(name: "test3.pow", scope: null, file: !6, line: 3, type: !7, scopeLine: 3, spFlags: DISPFlagLocalToUnit | DISPFlagDefinition, unit: !3, retainedNodes: !10)
!6 = !DIFile(filename: "isarmstrong.go", directory: ".")
!7 = !DISubroutineType(types: !8)
!8 = !{!9, !9, !9}
!9 = !DIBasicType(name: "int", size: 64, encoding: DW_ATE_signed)
!10 = !{}
!11 = !DILocalVariable(name: "a", arg: 1, scope: !5, file: !6, line: 3, type: !9)
!12 = !DILocation(line: 3, column: 10, scope: !5)
!13 = !DILocalVariable(name: "p", arg: 2, scope: !5, file: !6, line: 3, type: !9)
!14 = !DILocation(line: 3, column: 13, scope: !5)
!15 = !DILocalVariable(name: "$ret0", scope: !5, file: !6, line: 3, type: !9)
!16 = !DILocation(line: 3, column: 20, scope: !5)
!17 = !DILocation(line: 4, column: 7, scope: !18)
!18 = distinct !DILexicalBlock(scope: !5, file: !6, line: 3, column: 1)
!19 = !DILocalVariable(name: "r", scope: !18, file: !6, line: 4, type: !9)
!20 = !DILocation(line: 4, column: 2, scope: !5)
!21 = !DILocalVariable(name: "i", scope: !22, file: !6, line: 5, type: !9)
!22 = distinct !DILexicalBlock(scope: !18, file: !6, line: 5, column: 2)
!23 = !DILocation(line: 5, column: 6, scope: !18)
!24 = !DILocation(line: 5, column: 14, scope: !22)
!25 = !DILocation(line: 5, column: 18, scope: !22)
!26 = !DILocation(line: 5, column: 19, scope: !22)
!27 = !DILocation(line: 5, column: 16, scope: !22)
!28 = !DILocation(line: 6, column: 3, scope: !22)
!29 = !DILocation(line: 6, column: 8, scope: !22)
!30 = !DILocation(line: 6, column: 5, scope: !22)
!31 = !DILocation(line: 5, column: 23, scope: !22)
!32 = !DILocation(line: 8, column: 9, scope: !18)
!33 = !DILocation(line: 8, column: 2, scope: !18)
!34 = distinct !DISubprogram(name: "test3.IsArmstrong", scope: null, file: !6, line: 11, type: !35, scopeLine: 11, spFlags: DISPFlagDefinition, unit: !3, retainedNodes: !10)
!35 = !DISubroutineType(types: !36)
!36 = !{!37, !9}
!37 = !DIBasicType(name: "bool", size: 8, encoding: DW_ATE_unsigned)
!38 = !DILocalVariable(name: "n", arg: 1, scope: !34, file: !6, line: 11, type: !9)
!39 = !DILocation(line: 11, column: 18, scope: !34)
!40 = !DILocalVariable(name: "$ret1", scope: !34, file: !6, line: 11, type: !37)
!41 = !DILocation(line: 11, column: 25, scope: !34)
!42 = !DILocalVariable(name: "rightMost", scope: !43, file: !6, line: 12, type: !9)
!43 = distinct !DILexicalBlock(scope: !34, file: !6, line: 11, column: 1)
!44 = !DILocation(line: 12, column: 6, scope: !34)
!45 = !DILocalVariable(name: "sum", scope: !43, file: !6, line: 13, type: !9)
!46 = !DILocation(line: 13, column: 6, scope: !34)
!47 = !DILocation(line: 14, column: 20, scope: !43)
!48 = !DILocalVariable(name: "tempNum", scope: !43, file: !6, line: 14, type: !9)
!49 = !DILocation(line: 14, column: 6, scope: !34)
!50 = !DILocalVariable(name: "length", scope: !43, file: !6, line: 17, type: !9)
!51 = !DILocation(line: 17, column: 2, scope: !34)
!52 = !DILocation(line: 18, column: 9, scope: !43)
!53 = !DILocalVariable(name: "tmp", scope: !43, file: !6, line: 18, type: !9)
!54 = !DILocation(line: 18, column: 2, scope: !34)
!55 = !DILocation(line: 19, column: 6, scope: !43)
!56 = !DILocation(line: 19, column: 10, scope: !43)
!57 = !DILocation(line: 20, column: 3, scope: !43)
!58 = !DILocation(line: 20, column: 7, scope: !43)
!59 = !DILocation(line: 21, column: 3, scope: !43)
!60 = !DILocation(line: 25, column: 6, scope: !43)
!61 = !DILocation(line: 25, column: 14, scope: !43)
!62 = !DILocation(line: 26, column: 15, scope: !43)
!63 = !DILocation(line: 26, column: 23, scope: !43)
!64 = !DILocation(line: 26, column: 13, scope: !43)
!65 = !DILocation(line: 27, column: 14, scope: !43)
!66 = !DILocation(line: 27, column: 25, scope: !43)
!67 = !DILocation(line: 27, column: 10, scope: !43)
!68 = !DILocation(line: 27, column: 3, scope: !43)
!69 = !DILocation(line: 27, column: 7, scope: !43)
!70 = !DILocation(line: 30, column: 3, scope: !43)
!71 = !DILocation(line: 30, column: 11, scope: !43)
!72 = !DILocation(line: 33, column: 9, scope: !43)
!73 = !DILocation(line: 33, column: 14, scope: !43)
!74 = !DILocation(line: 33, column: 11, scope: !43)
!75 = !DILocation(line: 33, column: 2, scope: !43)

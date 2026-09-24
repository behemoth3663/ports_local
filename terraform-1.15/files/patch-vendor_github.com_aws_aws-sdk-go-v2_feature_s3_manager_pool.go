--- vendor/github.com/aws/aws-sdk-go-v2/feature/s3/manager/pool.go.orig	2026-03-26 18:11:26 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/feature/s3/manager/pool.go
@@ -151,7 +151,7 @@ func (p *maxSlicePool) ModifyCapacity(delta int) {
 	p.allocations = make(chan struct{}, p.max)
 
 	newAllocs := len(origAllocations) + delta
-	for range newAllocs {
+	for i := 0; i < newAllocs; i++ {
 		p.allocations <- struct{}{}
 	}
 

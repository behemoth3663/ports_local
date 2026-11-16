--- vendor/github.com/aws/aws-sdk-go-v2/service/chime/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/chime/serializers.go
@@ -11,7 +11,6 @@ import (
 	"github.com/aws/smithy-go/encoding/httpbinding"
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 )
 
@@ -25,10 +24,8 @@ func (m *awsRestjson1_serializeOpAssociatePhoneNumberW
 func (m *awsRestjson1_serializeOpAssociatePhoneNumberWithUser) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -77,7 +74,6 @@ func (m *awsRestjson1_serializeOpAssociatePhoneNumberW
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAssociatePhoneNumberWithUserInput(v *AssociatePhoneNumberWithUserInput, encoder *httpbinding.Encoder) error {
@@ -128,10 +124,8 @@ func (m *awsRestjson1_serializeOpAssociateSigninDelega
 func (m *awsRestjson1_serializeOpAssociateSigninDelegateGroupsWithAccount) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -180,7 +174,6 @@ func (m *awsRestjson1_serializeOpAssociateSigninDelega
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAssociateSigninDelegateGroupsWithAccountInput(v *AssociateSigninDelegateGroupsWithAccountInput, encoder *httpbinding.Encoder) error {
@@ -224,10 +217,8 @@ func (m *awsRestjson1_serializeOpBatchCreateRoomMember
 func (m *awsRestjson1_serializeOpBatchCreateRoomMembership) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -276,7 +267,6 @@ func (m *awsRestjson1_serializeOpBatchCreateRoomMember
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchCreateRoomMembershipInput(v *BatchCreateRoomMembershipInput, encoder *httpbinding.Encoder) error {
@@ -329,10 +319,8 @@ func (m *awsRestjson1_serializeOpBatchDeletePhoneNumbe
 func (m *awsRestjson1_serializeOpBatchDeletePhoneNumber) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -377,7 +365,6 @@ func (m *awsRestjson1_serializeOpBatchDeletePhoneNumbe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchDeletePhoneNumberInput(v *BatchDeletePhoneNumberInput, encoder *httpbinding.Encoder) error {
@@ -412,10 +399,8 @@ func (m *awsRestjson1_serializeOpBatchSuspendUser) Han
 func (m *awsRestjson1_serializeOpBatchSuspendUser) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -464,7 +449,6 @@ func (m *awsRestjson1_serializeOpBatchSuspendUser) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchSuspendUserInput(v *BatchSuspendUserInput, encoder *httpbinding.Encoder) error {
@@ -508,10 +492,8 @@ func (m *awsRestjson1_serializeOpBatchUnsuspendUser) H
 func (m *awsRestjson1_serializeOpBatchUnsuspendUser) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -560,7 +542,6 @@ func (m *awsRestjson1_serializeOpBatchUnsuspendUser) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchUnsuspendUserInput(v *BatchUnsuspendUserInput, encoder *httpbinding.Encoder) error {
@@ -604,10 +585,8 @@ func (m *awsRestjson1_serializeOpBatchUpdatePhoneNumbe
 func (m *awsRestjson1_serializeOpBatchUpdatePhoneNumber) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -652,7 +631,6 @@ func (m *awsRestjson1_serializeOpBatchUpdatePhoneNumbe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchUpdatePhoneNumberInput(v *BatchUpdatePhoneNumberInput, encoder *httpbinding.Encoder) error {
@@ -687,10 +665,8 @@ func (m *awsRestjson1_serializeOpBatchUpdateUser) Hand
 func (m *awsRestjson1_serializeOpBatchUpdateUser) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -739,7 +715,6 @@ func (m *awsRestjson1_serializeOpBatchUpdateUser) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsBatchUpdateUserInput(v *BatchUpdateUserInput, encoder *httpbinding.Encoder) error {
@@ -783,10 +758,8 @@ func (m *awsRestjson1_serializeOpCreateAccount) Handle
 func (m *awsRestjson1_serializeOpCreateAccount) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -831,7 +804,6 @@ func (m *awsRestjson1_serializeOpCreateAccount) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateAccountInput(v *CreateAccountInput, encoder *httpbinding.Encoder) error {
@@ -864,10 +836,8 @@ func (m *awsRestjson1_serializeOpCreateBot) HandleSeri
 func (m *awsRestjson1_serializeOpCreateBot) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -916,7 +886,6 @@ func (m *awsRestjson1_serializeOpCreateBot) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateBotInput(v *CreateBotInput, encoder *httpbinding.Encoder) error {
@@ -963,10 +932,8 @@ func (m *awsRestjson1_serializeOpCreateMeetingDialOut)
 func (m *awsRestjson1_serializeOpCreateMeetingDialOut) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1015,7 +982,6 @@ func (m *awsRestjson1_serializeOpCreateMeetingDialOut)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateMeetingDialOutInput(v *CreateMeetingDialOutInput, encoder *httpbinding.Encoder) error {
@@ -1067,10 +1033,8 @@ func (m *awsRestjson1_serializeOpCreatePhoneNumberOrde
 func (m *awsRestjson1_serializeOpCreatePhoneNumberOrder) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1115,7 +1079,6 @@ func (m *awsRestjson1_serializeOpCreatePhoneNumberOrde
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreatePhoneNumberOrderInput(v *CreatePhoneNumberOrderInput, encoder *httpbinding.Encoder) error {
@@ -1155,10 +1118,8 @@ func (m *awsRestjson1_serializeOpCreateRoom) HandleSer
 func (m *awsRestjson1_serializeOpCreateRoom) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1207,7 +1168,6 @@ func (m *awsRestjson1_serializeOpCreateRoom) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateRoomInput(v *CreateRoomInput, encoder *httpbinding.Encoder) error {
@@ -1254,10 +1214,8 @@ func (m *awsRestjson1_serializeOpCreateRoomMembership)
 func (m *awsRestjson1_serializeOpCreateRoomMembership) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1306,7 +1264,6 @@ func (m *awsRestjson1_serializeOpCreateRoomMembership)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateRoomMembershipInput(v *CreateRoomMembershipInput, encoder *httpbinding.Encoder) error {
@@ -1362,10 +1319,8 @@ func (m *awsRestjson1_serializeOpCreateUser) HandleSer
 func (m *awsRestjson1_serializeOpCreateUser) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1414,7 +1369,6 @@ func (m *awsRestjson1_serializeOpCreateUser) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateUserInput(v *CreateUserInput, encoder *httpbinding.Encoder) error {
@@ -1466,10 +1420,8 @@ func (m *awsRestjson1_serializeOpDeleteAccount) Handle
 func (m *awsRestjson1_serializeOpDeleteAccount) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1507,7 +1459,6 @@ func (m *awsRestjson1_serializeOpDeleteAccount) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteAccountInput(v *DeleteAccountInput, encoder *httpbinding.Encoder) error {
@@ -1537,10 +1488,8 @@ func (m *awsRestjson1_serializeOpDeleteEventsConfigura
 func (m *awsRestjson1_serializeOpDeleteEventsConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1578,7 +1527,6 @@ func (m *awsRestjson1_serializeOpDeleteEventsConfigura
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteEventsConfigurationInput(v *DeleteEventsConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -1617,10 +1565,8 @@ func (m *awsRestjson1_serializeOpDeletePhoneNumber) Ha
 func (m *awsRestjson1_serializeOpDeletePhoneNumber) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1658,7 +1604,6 @@ func (m *awsRestjson1_serializeOpDeletePhoneNumber) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeletePhoneNumberInput(v *DeletePhoneNumberInput, encoder *httpbinding.Encoder) error {
@@ -1688,10 +1633,8 @@ func (m *awsRestjson1_serializeOpDeleteRoom) HandleSer
 func (m *awsRestjson1_serializeOpDeleteRoom) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1729,7 +1672,6 @@ func (m *awsRestjson1_serializeOpDeleteRoom) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteRoomInput(v *DeleteRoomInput, encoder *httpbinding.Encoder) error {
@@ -1768,10 +1710,8 @@ func (m *awsRestjson1_serializeOpDeleteRoomMembership)
 func (m *awsRestjson1_serializeOpDeleteRoomMembership) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1809,7 +1749,6 @@ func (m *awsRestjson1_serializeOpDeleteRoomMembership)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteRoomMembershipInput(v *DeleteRoomMembershipInput, encoder *httpbinding.Encoder) error {
@@ -1857,10 +1796,8 @@ func (m *awsRestjson1_serializeOpDisassociatePhoneNumb
 func (m *awsRestjson1_serializeOpDisassociatePhoneNumberFromUser) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1898,7 +1835,6 @@ func (m *awsRestjson1_serializeOpDisassociatePhoneNumb
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisassociatePhoneNumberFromUserInput(v *DisassociatePhoneNumberFromUserInput, encoder *httpbinding.Encoder) error {
@@ -1937,10 +1873,8 @@ func (m *awsRestjson1_serializeOpDisassociateSigninDel
 func (m *awsRestjson1_serializeOpDisassociateSigninDelegateGroupsFromAccount) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1989,7 +1923,6 @@ func (m *awsRestjson1_serializeOpDisassociateSigninDel
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisassociateSigninDelegateGroupsFromAccountInput(v *DisassociateSigninDelegateGroupsFromAccountInput, encoder *httpbinding.Encoder) error {
@@ -2033,10 +1966,8 @@ func (m *awsRestjson1_serializeOpGetAccount) HandleSer
 func (m *awsRestjson1_serializeOpGetAccount) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2074,7 +2005,6 @@ func (m *awsRestjson1_serializeOpGetAccount) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetAccountInput(v *GetAccountInput, encoder *httpbinding.Encoder) error {
@@ -2104,10 +2034,8 @@ func (m *awsRestjson1_serializeOpGetAccountSettings) H
 func (m *awsRestjson1_serializeOpGetAccountSettings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2145,7 +2073,6 @@ func (m *awsRestjson1_serializeOpGetAccountSettings) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetAccountSettingsInput(v *GetAccountSettingsInput, encoder *httpbinding.Encoder) error {
@@ -2175,10 +2102,8 @@ func (m *awsRestjson1_serializeOpGetBot) HandleSeriali
 func (m *awsRestjson1_serializeOpGetBot) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2216,7 +2141,6 @@ func (m *awsRestjson1_serializeOpGetBot) HandleSeriali
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetBotInput(v *GetBotInput, encoder *httpbinding.Encoder) error {
@@ -2255,10 +2179,8 @@ func (m *awsRestjson1_serializeOpGetEventsConfiguratio
 func (m *awsRestjson1_serializeOpGetEventsConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2296,7 +2218,6 @@ func (m *awsRestjson1_serializeOpGetEventsConfiguratio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetEventsConfigurationInput(v *GetEventsConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -2335,10 +2256,8 @@ func (m *awsRestjson1_serializeOpGetGlobalSettings) Ha
 func (m *awsRestjson1_serializeOpGetGlobalSettings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2372,7 +2291,6 @@ func (m *awsRestjson1_serializeOpGetGlobalSettings) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetGlobalSettingsInput(v *GetGlobalSettingsInput, encoder *httpbinding.Encoder) error {
@@ -2393,10 +2311,8 @@ func (m *awsRestjson1_serializeOpGetPhoneNumber) Handl
 func (m *awsRestjson1_serializeOpGetPhoneNumber) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2434,7 +2350,6 @@ func (m *awsRestjson1_serializeOpGetPhoneNumber) Handl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetPhoneNumberInput(v *GetPhoneNumberInput, encoder *httpbinding.Encoder) error {
@@ -2464,10 +2379,8 @@ func (m *awsRestjson1_serializeOpGetPhoneNumberOrder) 
 func (m *awsRestjson1_serializeOpGetPhoneNumberOrder) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2505,7 +2418,6 @@ func (m *awsRestjson1_serializeOpGetPhoneNumberOrder) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetPhoneNumberOrderInput(v *GetPhoneNumberOrderInput, encoder *httpbinding.Encoder) error {
@@ -2535,10 +2447,8 @@ func (m *awsRestjson1_serializeOpGetPhoneNumberSetting
 func (m *awsRestjson1_serializeOpGetPhoneNumberSettings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2572,7 +2482,6 @@ func (m *awsRestjson1_serializeOpGetPhoneNumberSetting
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetPhoneNumberSettingsInput(v *GetPhoneNumberSettingsInput, encoder *httpbinding.Encoder) error {
@@ -2593,10 +2502,8 @@ func (m *awsRestjson1_serializeOpGetRetentionSettings)
 func (m *awsRestjson1_serializeOpGetRetentionSettings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2634,7 +2541,6 @@ func (m *awsRestjson1_serializeOpGetRetentionSettings)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetRetentionSettingsInput(v *GetRetentionSettingsInput, encoder *httpbinding.Encoder) error {
@@ -2664,10 +2570,8 @@ func (m *awsRestjson1_serializeOpGetRoom) HandleSerial
 func (m *awsRestjson1_serializeOpGetRoom) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2705,7 +2609,6 @@ func (m *awsRestjson1_serializeOpGetRoom) HandleSerial
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetRoomInput(v *GetRoomInput, encoder *httpbinding.Encoder) error {
@@ -2744,10 +2647,8 @@ func (m *awsRestjson1_serializeOpGetUser) HandleSerial
 func (m *awsRestjson1_serializeOpGetUser) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2785,7 +2686,6 @@ func (m *awsRestjson1_serializeOpGetUser) HandleSerial
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetUserInput(v *GetUserInput, encoder *httpbinding.Encoder) error {
@@ -2824,10 +2724,8 @@ func (m *awsRestjson1_serializeOpGetUserSettings) Hand
 func (m *awsRestjson1_serializeOpGetUserSettings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2865,7 +2763,6 @@ func (m *awsRestjson1_serializeOpGetUserSettings) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetUserSettingsInput(v *GetUserSettingsInput, encoder *httpbinding.Encoder) error {
@@ -2904,10 +2801,8 @@ func (m *awsRestjson1_serializeOpInviteUsers) HandleSe
 func (m *awsRestjson1_serializeOpInviteUsers) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2956,7 +2851,6 @@ func (m *awsRestjson1_serializeOpInviteUsers) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsInviteUsersInput(v *InviteUsersInput, encoder *httpbinding.Encoder) error {
@@ -3005,10 +2899,8 @@ func (m *awsRestjson1_serializeOpListAccounts) HandleS
 func (m *awsRestjson1_serializeOpListAccounts) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3046,7 +2938,6 @@ func (m *awsRestjson1_serializeOpListAccounts) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListAccountsInput(v *ListAccountsInput, encoder *httpbinding.Encoder) error {
@@ -3083,10 +2974,8 @@ func (m *awsRestjson1_serializeOpListBots) HandleSeria
 func (m *awsRestjson1_serializeOpListBots) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3124,7 +3013,6 @@ func (m *awsRestjson1_serializeOpListBots) HandleSeria
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListBotsInput(v *ListBotsInput, encoder *httpbinding.Encoder) error {
@@ -3162,10 +3050,8 @@ func (m *awsRestjson1_serializeOpListPhoneNumberOrders
 func (m *awsRestjson1_serializeOpListPhoneNumberOrders) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3203,7 +3089,6 @@ func (m *awsRestjson1_serializeOpListPhoneNumberOrders
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListPhoneNumberOrdersInput(v *ListPhoneNumberOrdersInput, encoder *httpbinding.Encoder) error {
@@ -3232,10 +3117,8 @@ func (m *awsRestjson1_serializeOpListPhoneNumbers) Han
 func (m *awsRestjson1_serializeOpListPhoneNumbers) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3273,7 +3156,6 @@ func (m *awsRestjson1_serializeOpListPhoneNumbers) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListPhoneNumbersInput(v *ListPhoneNumbersInput, encoder *httpbinding.Encoder) error {
@@ -3318,10 +3200,8 @@ func (m *awsRestjson1_serializeOpListRoomMemberships) 
 func (m *awsRestjson1_serializeOpListRoomMemberships) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3359,7 +3239,6 @@ func (m *awsRestjson1_serializeOpListRoomMemberships) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListRoomMembershipsInput(v *ListRoomMembershipsInput, encoder *httpbinding.Encoder) error {
@@ -3406,10 +3285,8 @@ func (m *awsRestjson1_serializeOpListRooms) HandleSeri
 func (m *awsRestjson1_serializeOpListRooms) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3447,7 +3324,6 @@ func (m *awsRestjson1_serializeOpListRooms) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListRoomsInput(v *ListRoomsInput, encoder *httpbinding.Encoder) error {
@@ -3489,10 +3365,8 @@ func (m *awsRestjson1_serializeOpListSupportedPhoneNum
 func (m *awsRestjson1_serializeOpListSupportedPhoneNumberCountries) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3530,7 +3404,6 @@ func (m *awsRestjson1_serializeOpListSupportedPhoneNum
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListSupportedPhoneNumberCountriesInput(v *ListSupportedPhoneNumberCountriesInput, encoder *httpbinding.Encoder) error {
@@ -3555,10 +3428,8 @@ func (m *awsRestjson1_serializeOpListUsers) HandleSeri
 func (m *awsRestjson1_serializeOpListUsers) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3596,7 +3467,6 @@ func (m *awsRestjson1_serializeOpListUsers) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListUsersInput(v *ListUsersInput, encoder *httpbinding.Encoder) error {
@@ -3642,10 +3512,8 @@ func (m *awsRestjson1_serializeOpLogoutUser) HandleSer
 func (m *awsRestjson1_serializeOpLogoutUser) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3683,7 +3551,6 @@ func (m *awsRestjson1_serializeOpLogoutUser) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsLogoutUserInput(v *LogoutUserInput, encoder *httpbinding.Encoder) error {
@@ -3722,10 +3589,8 @@ func (m *awsRestjson1_serializeOpPutEventsConfiguratio
 func (m *awsRestjson1_serializeOpPutEventsConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3774,7 +3639,6 @@ func (m *awsRestjson1_serializeOpPutEventsConfiguratio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutEventsConfigurationInput(v *PutEventsConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -3830,10 +3694,8 @@ func (m *awsRestjson1_serializeOpPutRetentionSettings)
 func (m *awsRestjson1_serializeOpPutRetentionSettings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3882,7 +3744,6 @@ func (m *awsRestjson1_serializeOpPutRetentionSettings)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutRetentionSettingsInput(v *PutRetentionSettingsInput, encoder *httpbinding.Encoder) error {
@@ -3926,10 +3787,8 @@ func (m *awsRestjson1_serializeOpRedactConversationMes
 func (m *awsRestjson1_serializeOpRedactConversationMessage) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3967,7 +3826,6 @@ func (m *awsRestjson1_serializeOpRedactConversationMes
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsRedactConversationMessageInput(v *RedactConversationMessageInput, encoder *httpbinding.Encoder) error {
@@ -4015,10 +3873,8 @@ func (m *awsRestjson1_serializeOpRedactRoomMessage) Ha
 func (m *awsRestjson1_serializeOpRedactRoomMessage) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4056,7 +3912,6 @@ func (m *awsRestjson1_serializeOpRedactRoomMessage) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsRedactRoomMessageInput(v *RedactRoomMessageInput, encoder *httpbinding.Encoder) error {
@@ -4104,10 +3959,8 @@ func (m *awsRestjson1_serializeOpRegenerateSecurityTok
 func (m *awsRestjson1_serializeOpRegenerateSecurityToken) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4145,7 +3998,6 @@ func (m *awsRestjson1_serializeOpRegenerateSecurityTok
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsRegenerateSecurityTokenInput(v *RegenerateSecurityTokenInput, encoder *httpbinding.Encoder) error {
@@ -4184,10 +4036,8 @@ func (m *awsRestjson1_serializeOpResetPersonalPIN) Han
 func (m *awsRestjson1_serializeOpResetPersonalPIN) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4225,7 +4075,6 @@ func (m *awsRestjson1_serializeOpResetPersonalPIN) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsResetPersonalPINInput(v *ResetPersonalPINInput, encoder *httpbinding.Encoder) error {
@@ -4264,10 +4113,8 @@ func (m *awsRestjson1_serializeOpRestorePhoneNumber) H
 func (m *awsRestjson1_serializeOpRestorePhoneNumber) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4305,7 +4152,6 @@ func (m *awsRestjson1_serializeOpRestorePhoneNumber) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsRestorePhoneNumberInput(v *RestorePhoneNumberInput, encoder *httpbinding.Encoder) error {
@@ -4335,10 +4181,8 @@ func (m *awsRestjson1_serializeOpSearchAvailablePhoneN
 func (m *awsRestjson1_serializeOpSearchAvailablePhoneNumbers) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4376,7 +4220,6 @@ func (m *awsRestjson1_serializeOpSearchAvailablePhoneN
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsSearchAvailablePhoneNumbersInput(v *SearchAvailablePhoneNumbersInput, encoder *httpbinding.Encoder) error {
@@ -4429,10 +4272,8 @@ func (m *awsRestjson1_serializeOpUpdateAccount) Handle
 func (m *awsRestjson1_serializeOpUpdateAccount) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4481,7 +4322,6 @@ func (m *awsRestjson1_serializeOpUpdateAccount) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateAccountInput(v *UpdateAccountInput, encoder *httpbinding.Encoder) error {
@@ -4528,10 +4368,8 @@ func (m *awsRestjson1_serializeOpUpdateAccountSettings
 func (m *awsRestjson1_serializeOpUpdateAccountSettings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4580,7 +4418,6 @@ func (m *awsRestjson1_serializeOpUpdateAccountSettings
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateAccountSettingsInput(v *UpdateAccountSettingsInput, encoder *httpbinding.Encoder) error {
@@ -4624,10 +4461,8 @@ func (m *awsRestjson1_serializeOpUpdateBot) HandleSeri
 func (m *awsRestjson1_serializeOpUpdateBot) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4676,7 +4511,6 @@ func (m *awsRestjson1_serializeOpUpdateBot) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateBotInput(v *UpdateBotInput, encoder *httpbinding.Encoder) error {
@@ -4727,10 +4561,8 @@ func (m *awsRestjson1_serializeOpUpdateGlobalSettings)
 func (m *awsRestjson1_serializeOpUpdateGlobalSettings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4775,7 +4607,6 @@ func (m *awsRestjson1_serializeOpUpdateGlobalSettings)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateGlobalSettingsInput(v *UpdateGlobalSettingsInput, encoder *httpbinding.Encoder) error {
@@ -4817,10 +4648,8 @@ func (m *awsRestjson1_serializeOpUpdatePhoneNumber) Ha
 func (m *awsRestjson1_serializeOpUpdatePhoneNumber) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4869,7 +4698,6 @@ func (m *awsRestjson1_serializeOpUpdatePhoneNumber) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdatePhoneNumberInput(v *UpdatePhoneNumberInput, encoder *httpbinding.Encoder) error {
@@ -4916,10 +4744,8 @@ func (m *awsRestjson1_serializeOpUpdatePhoneNumberSett
 func (m *awsRestjson1_serializeOpUpdatePhoneNumberSettings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4964,7 +4790,6 @@ func (m *awsRestjson1_serializeOpUpdatePhoneNumberSett
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdatePhoneNumberSettingsInput(v *UpdatePhoneNumberSettingsInput, encoder *httpbinding.Encoder) error {
@@ -4997,10 +4822,8 @@ func (m *awsRestjson1_serializeOpUpdateRoom) HandleSer
 func (m *awsRestjson1_serializeOpUpdateRoom) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5049,7 +4872,6 @@ func (m *awsRestjson1_serializeOpUpdateRoom) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateRoomInput(v *UpdateRoomInput, encoder *httpbinding.Encoder) error {
@@ -5100,10 +4922,8 @@ func (m *awsRestjson1_serializeOpUpdateRoomMembership)
 func (m *awsRestjson1_serializeOpUpdateRoomMembership) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5152,7 +4972,6 @@ func (m *awsRestjson1_serializeOpUpdateRoomMembership)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateRoomMembershipInput(v *UpdateRoomMembershipInput, encoder *httpbinding.Encoder) error {
@@ -5212,10 +5031,8 @@ func (m *awsRestjson1_serializeOpUpdateUser) HandleSer
 func (m *awsRestjson1_serializeOpUpdateUser) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5264,7 +5081,6 @@ func (m *awsRestjson1_serializeOpUpdateUser) HandleSer
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateUserInput(v *UpdateUserInput, encoder *httpbinding.Encoder) error {
@@ -5327,10 +5143,8 @@ func (m *awsRestjson1_serializeOpUpdateUserSettings) H
 func (m *awsRestjson1_serializeOpUpdateUserSettings) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -5379,7 +5193,6 @@ func (m *awsRestjson1_serializeOpUpdateUserSettings) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateUserSettingsInput(v *UpdateUserSettingsInput, encoder *httpbinding.Encoder) error {

PORTNAME=	f4
DISTVERSIONPREFIX=	v
DISTVERSION=	0.3.0-beta
CATEGORIES=	local misc

MAINTAINER=	none@FreeBSD.org
COMMENT=	Dual pane like a charm
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	BSD3CLAUSE

USES=		go:modules

USE_GITLAB=	nodefault

USE_GITHUB=	yes
GH_ACCOUNT=	unxed

GO_TARGET=	./cmd/${PORTNAME}

CGO_ENABLED=	0
GO_BUILDFLAGS+=	-gcflags=github.com/go-webgpu/goffi/internal/fakecgo=-std
GO_LDFLAGS+=	-w
GO_LDFLAGS+=	-X github.com/${GH_ACCOUNT}/${GH_PROJECT}/internal/app.buildVersion=${DISTVERSIONFULL}
GO_ENV+=	GOEXPERIMENT=norandomizedheapbase64

OPTIONS_DEFINE=	LITE
OPTIONS_DEFAULT=

LITE_VARS=	GO_BUILDFLAGS+=-tags=lite
LITE_VARS+=	GO_LDFLAGS+=-buildid=

M2T_OFFLINE=	true
.export	M2T_OFFLINE

post-extract: .PHONY
	${TEST} -f ${WRKSRC}/vendor/modules.txt || ${TEST} ! -f ${PATCHDIR}/modules.txt || ${CP} ${PATCHDIR}/modules.txt ${WRKSRC}/vendor
	${TEST} -z ${WRKSRC_unxed_libwinescape} || ${MV} ${WRKDIR}/${WRKSRC_unxed_libwinescape:C,.*/lib,,1} ${WRKSRC}/vendor/github.com/unxed/libwinescape

post-install: .PHONY
	${MKDIR} ${STAGEDIR}${DATADIR}/help
	${INSTALL_DATA} ${WRKSRC}/internal/dialog/help/*.hlf ${STAGEDIR}${DATADIR}/help
	${INSTALL_DATA} ${WRKSRC}/internal/dialog/help/*.md ${STAGEDIR}${DATADIR}/help
	${MKDIR} ${STAGEDIR}${DATADIR}/lang
	${INSTALL_DATA} ${WRKSRC}/internal/i18n/lang/*.lng ${STAGEDIR}${DATADIR}/lang
	${INSTALL_DATA} ${WRKSRC}/internal/i18n/lang/*.md ${STAGEDIR}${DATADIR}/lang
	${INSTALL_DATA} ${WRKSRC}/f4.example.ini ${STAGEDIR}${DATADIR}

pre-deps: .PHONY
	${TEST} -d ${PATCHDIR}_ && ${RM} -r ${PATCHDIR} || ${MV} ${PATCHDIR} ${PATCHDIR}_

pre-deps-all: .PHONY
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}_
	${RM} ${.CURDIR}/generated.deps.mk
	( ${MAKE} -C ${.CURDIR} make-from-deps GHT_ACNT=SSLMate GHT_PROJ=go-pkcs12 GHT_PATH=software.sslmate.com/src/go-pkcs12 )

make-from-deps: .PHONY
	@${SED} -rn \
		's|^[[:blank:]]*#[[:blank:]]*::([v0-9.a-f]+):.*/(vendor/${GHT_PATH})[[:blank:]].*$$|GH_TUPLE+=${GHT_ACNT}:${GHT_PROJ}:\1:${GHT_ACNT:tl:S/-/_/g}_${GHT_PROJ:tl:S/-/_/g}/\2|p' \
	${.CURDIR}/Makefile.deps >>${.CURDIR}/generated.deps.mk

post-deps-all: .PHONY
	${TEST} ! -d ${PATCHDIR}_ || ${MV} ${PATCHDIR}_ ${PATCHDIR}

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"

#-aws:aws-sdk-go-v2:v0.3.15:aws_aws_sdk_go_v2_transfermanager/vendor/github.com/aws/aws-sdk-go-v2/feature/s3/transfermanager
#-aws:aws-sdk-go-v2:v1.107.3:aws_aws_sdk_go_v2_s3/vendor/github.com/aws/aws-sdk-go-v2/service/s3
#-aws:aws-sdk-go-v2:v1.13.17:aws_aws_sdk_go_v2_accept_encoding/vendor/github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding
#-aws:aws-sdk-go-v2:v1.13.38:aws_aws_sdk_go_v2_presigned_url/vendor/github.com/aws/aws-sdk-go-v2/service/internal/presigned-url
#-aws:aws-sdk-go-v2:v1.18.38:aws_aws_sdk_go_v2_imds/vendor/github.com/aws/aws-sdk-go-v2/feature/ec2/imds
#-aws:aws-sdk-go-v2:v1.19.37:aws_aws_sdk_go_v2_credentials/vendor/github.com/aws/aws-sdk-go-v2/credentials
#-aws:aws-sdk-go-v2:v1.19.39:aws_aws_sdk_go_v2_s3shared/vendor/github.com/aws/aws-sdk-go-v2/service/internal/s3shared
#-aws:aws-sdk-go-v2:v1.22.42:aws_aws_sdk_go_v2_manager/vendor/github.com/aws/aws-sdk-go-v2/feature/s3/manager
#-aws:aws-sdk-go-v2:v1.32.38:aws_aws_sdk_go_v2_config/vendor/github.com/aws/aws-sdk-go-v2/config
#-aws:aws-sdk-go-v2:v1.33.7:aws_aws_sdk_go_v2_sso/vendor/github.com/aws/aws-sdk-go-v2/service/sso
#-aws:aws-sdk-go-v2:v1.38.7:aws_aws_sdk_go_v2_ssooidc/vendor/github.com/aws/aws-sdk-go-v2/service/ssooidc
#-aws:aws-sdk-go-v2:v1.4.38:aws_aws_sdk_go_v2_configsources/vendor/github.com/aws/aws-sdk-go-v2/internal/configsources
#-aws:aws-sdk-go-v2:v1.4.39:aws_aws_sdk_go_v2_v4a/vendor/github.com/aws/aws-sdk-go-v2/internal/v4a
# aws:aws-sdk-go-v2:v1.43.7:aws_aws_sdk_go_v2/vendor/github.com/aws/aws-sdk-go-v2
#-aws:aws-sdk-go-v2:v1.45.7:aws_aws_sdk_go_v2_sts/vendor/github.com/aws/aws-sdk-go-v2/service/sts
#-aws:aws-sdk-go-v2:v1.5.7:aws_aws_sdk_go_v2_signin/vendor/github.com/aws/aws-sdk-go-v2/service/signin
#-aws:aws-sdk-go-v2:v1.7.18:aws_aws_sdk_go_v2_eventstream/vendor/github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream
#-aws:aws-sdk-go-v2:v1.9.31:aws_aws_sdk_go_v2_checksum/vendor/github.com/aws/aws-sdk-go-v2/service/internal/checksum
#-aws:aws-sdk-go-v2:v2.7.38:aws_aws_sdk_go_v2_v2/vendor/github.com/aws/aws-sdk-go-v2/internal/endpoints/v2
GH_TUPLE:=	${GH_TUPLE:Naws*aws_aws_sdk_go_v2_*/vendor/github.com/aws/aws-sdk-go-v2/*}

#-googleapis:google-cloud-go:v0.18.2:googleapis_google_cloud_go_auth/vendor/cloud.google.com/go/auth
#+googleapis:google-cloud-go:auth/v0.18.2:googleapis_google_cloud_go_auth/vendor/cloud.google.com/go
GH_TUPLE:=	${GH_TUPLE:C,:([^:]+:[^:/]+/.*/cloud.google.com/go)/(auth),:\2/\1,1}

#-unxed:libwinescape:v0.2.1:unxed_libwinescape/vendor/github.com/unxed/libwinescape
#+unxed:libwinescape:v0.2.1:unxed_libwinescape
GH_TUPLE:=	${GH_TUPLE:S,/vendor/github.com/unxed/libwinescape,,1}

#-googleapis:gax-go:v2.11.0:googleapis_gax_go_v2/vendor/github.com/googleapis/gax-go/v2
#+googleapis:gax-go:v2.11.0:googleapis_gax_go_v2/vendor/github.com/googleapis/gax-go
GH_TUPLE:=	${GH_TUPLE:S|/gax-go/v2|/gax-go|1}

#-google:go-genproto:3dc84a4a5aaa:google_go_genproto_rpc/vendor/google.golang.org/genproto/googleapis/rpc
#+google:go-genproto:3dc84a4a5aaa:google_go_genproto_rpc/vendor/google.golang.org/genproto/googleapis
GH_TUPLE:=	${GH_TUPLE:S|/googleapis/rpc||1}

.sinclude "${.CURDIR}/generated.deps.mk"
.include <bsd.port.mk>
		  	
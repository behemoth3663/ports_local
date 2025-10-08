PORTNAME=	ssosync
PORTVERSION=	2.3.3
DISTVERSIONPREFIX=	v
#PORTREVISION=	1
CATEGORIES=	local sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	Populate AWS SSO directly with Google Apps users
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	APACHE20

BUILD_DEPENDS=	git:devel/git

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	awslabs

PLIST_FILES=	bin/${PORTNAME}

M2T_OFFLINE=	true
.export	M2T_OFFLINE

GIT_COMMIT=	git ls-remote \
		https://github.com/${GH_ACCOUNT}/${PORTNAME} \
		refs/tags/${DISTVERSIONPREFIX}${PORTVERSION} | \
		${SED} -rn 's/(.{7}).*/\1/p'

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags -static \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/cmd.version=${PORTVERSION}' \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/cmd.commit=${GIT_COMMIT:sh}' \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/cmd.date=${%FT%TZ:L:gmtime}' \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/cmd.builtBy=${OPSYS}'"

post-extract: .PHONY
	@${CP} ${PATCHDIR}/modules.txt ${WRKSRC}/vendor
	@${FIND} ${WRKSRC} -name '*_test.go' -type f -delete

pre-deps: .PHONY
	test -d ${PATCHDIR}_ && ${RM} -r ${PATCHDIR} || ${MV} ${PATCHDIR} ${PATCHDIR}_

pre-deps-all: .PHONY
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}_

post-deps-all: .PHONY
	test ! -d ${PATCHDIR}_ || ${MV} ${PATCHDIR}_ ${PATCHDIR}

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"

#-aws:aws-sdk-go-v2:v1.3.15:aws_aws_sdk_go_v2_configsources/vendor/github.com/aws/aws-sdk-go-v2/internal/configsources
# aws:aws-sdk-go-v2:v1.30.3:aws_aws_sdk_go_v2/vendor/github.com/aws/aws-sdk-go-v2
#-aws:aws-sdk-go-v2:v1.30.3:aws_aws_sdk_go_v2_sts/vendor/github.com/aws/aws-sdk-go-v2/service/sts
GH_TUPLE:=	${GH_TUPLE:Naws*aws_aws_sdk_go_v2_*/vendor/github.com/aws/aws-sdk-go-v2/*}

GH_TUPLE:=	${GH_TUPLE:S|/gax-go/v2|/gax-go|1}

#-googleapis:google-cloud-go:v0.16.4:googleapis_google_cloud_go_auth/vendor/cloud.google.com/go/auth
#+googleapis:google-cloud-go:auth/v0.16.4:googleapis_google_cloud_go_auth/vendor/cloud.google.com/go
GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)([^:]+[^/]+.*/cloud.google.com/go)/(auth)|\1\3/\2|}

#-google:go-genproto:a7a43d27e69b:google_go_genproto/vendor/google.golang.org/genproto/googleapis/rpc
#+google:go-genproto:a7a43d27e69b:google_go_genproto/vendor/google.golang.org/genproto/googleapis
GH_TUPLE:=	${GH_TUPLE:S|/googleapis/rpc||1}

.include <bsd.port.mk>

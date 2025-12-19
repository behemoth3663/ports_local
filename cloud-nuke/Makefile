PORTNAME=	cloud-nuke
PORTVERSION=	0.46.0
DISTVERSIONPREFIX=	v
#PORTREVISION=	1
CATEGORIES=	local sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	CLI tool to delete all resources in an AWS account
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE.txt

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	gruntwork-io

PLIST_FILES=	bin/${PORTNAME}

CGO_ENABLED=	0
GO_BUILDFLAGS+=	-ldflags="-s -w \
	-X 'main.VERSION=${DISTVERSIONPREFIX}${PORTVERSION}'"

#GO_ENV+=	GOEXPERIMENT=greenteagc

M2T_OFFLINE=	true
.export	M2T_OFFLINE

_OPTIONS_extract=	600:pre-post-extract
#Resolve the error:
#mv: rename /usr/obj/usr/ports/local/cloud-nuke/work/gruntwork-cli-0.7.0 to /usr/obj/usr/ports/local/cloud-nuke/work/cloud-nuke-0.29.4/vendor/github.com/gruntwork-io/gruntwork-cli: No such file or directory
#/usr/obj/usr/ports/local/cloud-nuke/work/gruntwork-cli-0.7.0 -> /usr/obj/usr/ports/local/cloud-nuke/work/go-commons-0.7.0
#/usr/obj/usr/ports/local/cloud-nuke/work/gruntwork-cli-0.16.1 -> /usr/obj/usr/ports/local/cloud-nuke/work/go-commons-0.16.1
pre-post-extract: .PHONY .SILENT
	test -z '${WRKDIR}' || ${FIND} ${WRKDIR} -maxdepth 1 -type d -name 'go-commons-*' | while read src; do \
		${LN} -s "$${src}" "$${src%%go-commons-*}gruntwork-cli-$${src##*go-commons-}"; \
	done

post-extract: .PHONY
	@${CP} ${PATCHDIR}/modules.txt ${WRKSRC}/vendor
	@${RM} -r ${WRKSRC}/vendor/github.com/aws/aws-sdk-go-v2/service/securityhub
	@test -z ${WRKSRC_aws_aws_sdk_go_v2_securityhub} || ${RLN} ${WRKSRC_aws_aws_sdk_go_v2_securityhub}/service/securityhub ${WRKSRC}/vendor/github.com/aws/aws-sdk-go-v2/service

pre-deps: .PHONY
	test -d ${PATCHDIR}_ && ${RM} -r ${PATCHDIR} || ${MV} ${PATCHDIR} ${PATCHDIR}_

pre-deps-all: .PHONY
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}_
	${RM} ${.CURDIR}/generated.deps.mk
	( make -C ${.CURDIR} make-from-deps GHT_ACNT=atomicgo GHT_PROJ=cursor GHT_PATH=atomicgo.dev/cursor )
	( make -C ${.CURDIR} make-from-deps GHT_ACNT=atomicgo GHT_PROJ=keyboard GHT_PATH=atomicgo.dev/keyboard )

make-from-deps: .PHONY
	@${SED} -rn \
		's|^[[:space:]]*#[[:space:]]*::([v0-9.a-f]+):.*/(vendor/${GHT_PATH})[[:space:]].*$$|GH_TUPLE+=${GHT_ACNT}:${GHT_PROJ}:\1:${GHT_ACNT:tl:S/-/_/g}_${GHT_PROJ:tl:S/-/_/g}/\2|p' \
	${.CURDIR}/Makefile.deps >>${.CURDIR}/generated.deps.mk

post-deps-all: .PHONY
	test -d ${PATCHDIR}_ && ${MV} ${PATCHDIR}_ ${PATCHDIR} || :

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"

#-aws:aws-sdk-go-v2:v1.55.8:aws_aws_sdk_go_v2_securityhub/vendor/github.com/aws/aws-sdk-go-v2/service/securityhub
#+aws:aws-sdk-go-v2:service/securityhub/v1.55.8:aws_aws_sdk_go_v2_securityhub
GH_TUPLE:=	${GH_TUPLE:C,^([^:]+:[^:]+:)([^:]+:[^/]+).*/(service/securityhub)$,\1\3/\2,1}

#-		aws:aws-sdk-go-v2:v1.3.15:aws_aws_sdk_go_v2_configsources/vendor/github.com/aws/aws-sdk-go-v2/internal/configsources
# 		aws:aws-sdk-go-v2:v1.30.3:aws_aws_sdk_go_v2/vendor/github.com/aws/aws-sdk-go-v2
#-		aws:aws-sdk-go-v2:v1.30.3:aws_aws_sdk_go_v2_sts/vendor/github.com/aws/aws-sdk-go-v2/service/sts
GH_TUPLE:=	${GH_TUPLE:Naws*aws_aws_sdk_go_v2_*/vendor/github.com/aws/aws-sdk-go-v2/*}

# googleapis:google-cloud-go:v0.112.1:googleapis_google_cloud_go/vendor/cloud.google.com/go
#-googleapis:google-cloud-go:v1.1.6:googleapis_google_cloud_go_iam/vendor/cloud.google.com/go/iam
#-googleapis:google-cloud-go:v1.39.1:googleapis_google_cloud_go_storage/vendor/cloud.google.com/go/storage
GH_TUPLE:=	${GH_TUPLE:Ngoogleapis\:google-cloud-go\:*/cloud.google.com/go/*}

#-googleapis:gax-go:v2.13.0:googleapis_gax_go_v2/vendor/github.com/googleapis/gax-go/v2
#+googleapis:gax-go:v2.13.0:googleapis_gax_go_v2/vendor/github.com/googleapis/gax-go
GH_TUPLE:=	${GH_TUPLE:S|/gax-go/v2|/gax-go|1}

#-census-instrumentation:opencensus-go:v0.24.0:census_instrumentation_opencensus_go/vendor/go.opencensus.io
GH_TUPLE:=	${GH_TUPLE:Ncensus-instrumentation\:opencensus-go\:*}

.sinclude "${.CURDIR}/generated.deps.mk"
.include <bsd.port.mk>

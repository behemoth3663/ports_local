PORTNAME=	cloud-nuke
PORTVERSION=	0.38.1
DISTVERSIONPREFIX=	v
PORTREVISION=	2
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

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags -static \
	-X 'main.VERSION=${DISTVERSIONPREFIX}${PORTVERSION}'"

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

pre-deps post-deps-all: .PHONY
pre-deps-all: .PHONY
	${RM} ${.CURDIR}/generated.deps.mk
	( make -C ${.CURDIR} make-from-deps GHT_ACNT=atomicgo GHT_PROJ=cursor GHT_PATH=atomicgo.dev/cursor )
	( make -C ${.CURDIR} make-from-deps GHT_ACNT=atomicgo GHT_PROJ=keyboard GHT_PATH=atomicgo.dev/keyboard )

make-from-deps: .PHONY
	@${SED} -rn \
		's|^[[:space:]]*#[[:space:]]*::(v[0-9.]+):.*/(vendor/${GHT_PATH})[[:space:]].*$$|GH_TUPLE+=${GHT_ACNT}:${GHT_PROJ}:\1:${GHT_ACNT:tl:S/-/_/}_${GHT_PROJ:tl:S/-/_/}/\2|p' \
	${.CURDIR}/Makefile.deps >>${.CURDIR}/generated.deps.mk

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"

#-		aws:aws-sdk-go-v2:v1.3.15:aws_aws_sdk_go_v2_configsources/vendor/github.com/aws/aws-sdk-go-v2/internal/configsources
# 		aws:aws-sdk-go-v2:v1.30.3:aws_aws_sdk_go_v2/vendor/github.com/aws/aws-sdk-go-v2
#-		aws:aws-sdk-go-v2:v1.30.3:aws_aws_sdk_go_v2_sts/vendor/github.com/aws/aws-sdk-go-v2/service/sts
GH_TUPLE:=	${GH_TUPLE:Naws*aws_aws_sdk_go_v2_*/vendor/github.com/aws/aws-sdk-go-v2/*}

.sinclude "${.CURDIR}/generated.deps.mk"
.include <bsd.port.mk>

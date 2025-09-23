PORTNAME=	aws-ec2-instance-selector
PORTVERSION=	3.1.2
DISTVERSIONPREFIX=	v
#PORTREVISION=	1
CATEGORIES=	local sysutils
DISTNAME=	amazon-ec2-instance-selector

MAINTAINER=	none@freebsd.org
COMMENT=	CLI tool which recommends instance types based on resource criteria
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	APACHE20

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	aws
GH_PROJECT=	${DISTNAME}

GO_TARGET=	./cmd/main.go:${PORTNAME}

PLIST_FILES=	bin/${PORTNAME}

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags -static \
	-X 'main.versionID=${DISTVERSIONPREFIX}${PORTVERSION}'"

M2T_OFFLINE=	true
.export	M2T_OFFLINE

#pre-build: .PHONY
#	${FIND} ${WRKSRC} -name go.mod -type f -execdir ${GO_CMD} mod edit -go=1.25 \;

post-extract: .PHONY
	${CP} ${PATCHDIR}/modules.txt ${WRKSRC}/vendor
.if defined(WRKSRC_charmbracelet_x_term)
	${RM} -r ${WRKSRC}/vendor/github.com/charmbracelet/x/term
	${RLN} ${WRKSRC_charmbracelet_x_term}/term ${WRKSRC}/vendor/github.com/charmbracelet/x/term
.endif

pre-deps post-deps-all: .PHONY

pre-deps-all: .PHONY
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}
	${RM} ${.CURDIR}/generated.deps.mk
	( make -C ${.CURDIR} make-from-deps GHT_ACNT=darccio GHT_PROJ=mergo GHT_PATH=dario.cat/mergo )

make-from-deps: .PHONY
	@${SED} -rn \
		's|^[[:space:]]*#[[:space:]]*::([v0-9.a-f]+):.*/(vendor/${GHT_PATH})[[:space:]].*$$|GH_TUPLE+=${GHT_ACNT}:${GHT_PROJ}:\1:${GHT_ACNT:tl:S/-/_/g}_${GHT_PROJ:tl:S/-/_/g}/\2|p' \
	${.CURDIR}/Makefile.deps >>${.CURDIR}/generated.deps.mk

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"

#-aws:aws-sdk-go-v2:v1.3.15:aws_aws_sdk_go_v2_configsources/vendor/github.com/aws/aws-sdk-go-v2/internal/configsources
# aws:aws-sdk-go-v2:v1.30.3:aws_aws_sdk_go_v2/vendor/github.com/aws/aws-sdk-go-v2
#-aws:aws-sdk-go-v2:v1.30.3:aws_aws_sdk_go_v2_sts/vendor/github.com/aws/aws-sdk-go-v2/service/sts
GH_TUPLE:=	${GH_TUPLE:Naws*aws_aws_sdk_go_v2_*/vendor/github.com/aws/aws-sdk-go-v2/*}

#-blang:semver:v4.0.0:blang_semver_v4/vendor/github.com/blang/semver/v4
#+blang:semver:v4.0.0:blang_semver_v4/vendor/github.com/blang/semver
GH_TUPLE:=	${GH_TUPLE:S|/semver/v4|/semver|1}

#-charmbracelet:x:     v0.8.0:charmbracelet_x_ansi/vendor/github.com/charmbracelet/x/ansi
#+charmbracelet:x:ansi/v0.8.0:charmbracelet_x_ansi/vendor/github.com/charmbracelet/x
GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)(.*/charmbracelet/x)/(ansi)|\1\3/\2|1}

#-charmbracelet:x:     v0.2.1:charmbracelet_x_term/vendor/github.com/charmbracelet/x/term
#+charmbracelet:x:term/v0.2.1:charmbracelet_x_term
GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)([^:]+[^/]+).*/charmbracelet/x/(term)|\1\3/\2|1}

.sinclude "${.CURDIR}/generated.deps.mk"
.include <bsd.port.mk>

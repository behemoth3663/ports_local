PORTNAME=	aws-nuke
PORTVERSION=	2.24.1
DISTVERSIONPREFIX=	v
CATEGORIES=	local sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	Remove all resources from an AWS account
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

BUILD_DEPENDS=	git:devel/git

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	rebuy-de

PLIST_FILES=	bin/${PORTNAME}

GIT_COMMIT=	git ls-remote \
		https://github.com/${GH_ACCOUNT}/${GH_PROJECT} \
		refs/tags/${DISTVERSIONPREFIX}${PORTVERSION} | \
		${SED} -rn 's/(.{7}).*/\1/p'

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags '-static' \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/v2/cmd.BuildVersion=${PORTVERSION}' \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/v2/cmd.BuildDate=${%FT%TZ:L:gmtime}' \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/v2/cmd.BuildHash=$$(${GIT_COMMIT})' \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/v2/cmd.BuildEnvironment='"

pre-deps: .PHONY
	@cd ${WRKSRC} && ${SETENV} GOPATH=${WRKDIR}/.gopath ${GO_CMD} generate ./...

pre-deps-all post-deps-all: .PHONY

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>

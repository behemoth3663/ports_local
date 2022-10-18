#

PORTNAME=	aws-nuke
PORTVERSION=	2.20.0
DISTVERSIONPREFIX=	v
CATEGORIES=	local sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	Remove all resources from an AWS account
WWW=		https://github.com/rebuy-de/aws-nuke

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

BUILD_DEPENDS=	git:devel/git

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	rebuy-de

PLIST_FILES=	bin/${PORTNAME}

GH_COMMIT=	git ls-remote --tags \
		https://github.com/${GH_ACCOUNT}/${PORTNAME} \
		refs/tags/${DISTVERSIONPREFIX}${PORTVERSION}

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags '-static' \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/cmd.BuildVersion=${DISTVERSIONPREFIX}${PORTVERSION}' \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/cmd.BuildDate=${%FT%TZ:L:gmtime}' \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/cmd.BuildHash=${:!${GH_COMMIT}!:[1]:C/(.{7}).*/\1/}' \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/cmd.BuildEnvironment='"

post-patch: .PHONY
.if ${.TARGETS:Mgomod-vendor}
	@cd ${WRKSRC} && ${SETENV} GOPATH=${WRKDIR}/.gopath ${GO_CMD} generate ./...
.endif

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>

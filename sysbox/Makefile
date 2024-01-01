PORTNAME=	sysbox
PORTVERSION=	0.18.0
DISTVERSIONPREFIX=	release-
CATEGORIES=	local sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	sysadmin/scripting utilities, distributed as a single binary
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	GPLv2
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	skx

PLIST_FILES=	bin/${PORTNAME}

JQ_GETVER_EXPR=	[.[].name | gsub("^[a-z]+-"; "") | select(test("^[0-9.]+$$"; "n"))][0]

#post-extract: .PHONY
#	${REINPLACE_CMD} -r \
#		-e 's|^([[:space:]]*go[[:space:]]+).*$$|\11.21|' \
#	${WRKSRC}/go.mod
#	cd ${WRKSRC} && ${SETENV} GOPATH=${WRKDIR}/.gopath ${GO_CMD} mod tidy -v

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags '-static' \
	-X 'main.version=${PORTVERSION}'"

pre-deps pre-deps-all post-deps-all: .PHONY

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>

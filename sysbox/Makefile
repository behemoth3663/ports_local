#

PORTNAME=	sysbox
PORTVERSION=	0.16.0
DISTVERSIONPREFIX=	release-
CATEGORIES=	local sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	sysadmin/scripting utilities, distributed as a single binary

LICENSE=	GPLv2
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	skx

PLIST_FILES=	bin/${PORTNAME}

GO_BUILDFLAGS+=	-ldflags="-s -w -X 'main.version=${PORTVERSION}'"

JQ_GETVER_EXPR=	[.[].name | gsub("^[a-z]+-"; "") | select(test("^[0-9.]+$$"; "n"))][0]

fetch-version: .PHONY get-version

#post-extract: .PHONY
#	${REINPLACE_CMD} -r \
#		-e 's|^([[:space:]]*go[[:space:]]+).*$$|\11.18|' \
#	${WRKSRC}/go.mod
#	cd ${WRKSRC} && ${SETENV} GOPATH=${WRKDIR}/.gopath ${GO_CMD} mod tidy -v

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>

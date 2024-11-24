PORTNAME=	sysbox
PORTVERSION=	0.19.0
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

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags '-static' \
	-X 'main.version=${PORTVERSION}'"

post-extract: .PHONY
	test -f ${WRKSRC}/vendor/modules.txt || test ! -f ${PATCHDIR}/modules.txt || ${CP} ${PATCHDIR}/modules.txt ${WRKSRC}/vendor

pre-deps post-deps-all: .PHONY
pre-deps-all: .PHONY
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"

#-google:go-genproto:e639e219e697:google_go_genproto/vendor/google.golang.org/genproto/googleapis/rpc
#+google:go-genproto:782d3b101e98:google_go_genproto/vendor/google.golang.org/genproto
GH_TUPLE:=	${GH_TUPLE:S|/googleapis/rpc$||}

.include <bsd.port.mk>

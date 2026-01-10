PORTNAME=	checkmake
PORTVERSION=	0.3.1
DISTVERSIONPREFIX=	v
CATEGORIES=	local devel

MAINTAINER=	none@freebsd.org
COMMENT=	Linter/analyzer for Makefiles
WWW=		https://github.com/mrtazz/checkmake

LICENSE=	MIT

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	mrtazz

PLIST_FILES=	bin/${PORTNAME}

GO_TARGET=	./cmd/${PORTNAME}

CGO_ENABLED=	0
GO_BUILDFLAGS+=	-ldflags="-s -w \
	-X 'main.version=${PORTVERSION}' \
	-X 'main.buildTime=${%FT%TZ:L:gmtime}' \
	-X 'main.builder=FreeBSD ports' \
	-X 'main.goversion=go${GO_VERSION}'"

OPTIONS_DEFINE=	MANPAGES
OPTIONS_DEFAULT=MANPAGES

MANPAGES_BUILD_DEPENDS=	pandoc:textproc/hs-pandoc
MANPAGES_PLIST_FILES=	share/man/man1/${PORTNAME}.1.gz

post-build-MANPAGES-on: .PHONY
	# Retrieve buildTime from binary built from patched sources
	${SED} 's/REPLACE_DATE/${:!${GO_WRKDIR_BIN}/${GO_TARGET:T} --version!:[5]}/' ${WRKSRC}/man/man1/${PORTNAME}.1.md | \
		pandoc --standalone --to=man --output=${WRKSRC}/man/man1/${PORTNAME}.1

post-stage-MANPAGES-on: .PHONY
	${INSTALL_MAN} ${WRKSRC}/man/man1/${PORTNAME}.1 ${STAGEDIR}${LOCALBASE}/share/man/man1

post-extract: .PHONY
	@${CP} ${PATCHDIR}/modules.txt ${WRKSRC}/vendor

pre-deps: .PHONY
	test -d ${PATCHDIR}_ && ${RM} -r ${PATCHDIR} || ${MV} ${PATCHDIR} ${PATCHDIR}_

pre-deps-all: .PHONY
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}_

post-deps-all: .PHONY
	test ! -d ${PATCHDIR}_ || ${MV} ${PATCHDIR}_ ${PATCHDIR}

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>

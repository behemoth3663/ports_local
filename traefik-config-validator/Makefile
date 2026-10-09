PORTNAME=	traefik-config-validator
PORTVERSION=	0.0.2
DISTVERSIONPREFIX=	v
PORTREVISION=	2
CATEGORIES=	local sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	Validate your Traefik configuration files
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	APACHE20

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	otto-de

GO_TARGET=	./cmd/${PORTNAME}

CGO_ENABLED=	0
GO_LDFLAGS+=	-w
GO_LDFLAGS+=	-X main.version=${PORTVERSION}
GO_ENV+=	GOEXPERIMENT=norandomizedheapbase64

PLIST_FILES=	bin/${PORTNAME}

M2T_OFFLINE=	true
.export	M2T_OFFLINE

post-extract: .PHONY
	@${MKDIR} ${WRKSRC}/vendor
	@${CP} ${PATCHDIR}/modules.txt ${WRKSRC}/vendor

pre-deps: .PHONY
	${TEST} -d ${PATCHDIR}_ && ${RM} -r ${PATCHDIR} || ${MV} ${PATCHDIR} ${PATCHDIR}_
	${CP} -a ${WRKSRC}/go.mod ${WRKSRC}/go.mod.orig
	${CP} -a ${WRKSRC}/go.sum ${WRKSRC}/go.sum.orig
	cd ${WRKSRC}/${GO_TARGET} && ${GO_CMD} get -u
	cd ${WRKSRC}/${GO_TARGET} && ${GO_CMD} mod tidy
	cd ${WRKSRC} && ${SETENV} TZ=UTC ${DIFF} -u --minimal -- go.mod.orig go.mod | \
		${SED} -e '/^---/s|\.[0-9]* +0000$$| UTC|' -e '/^+++/s|\([[:blank:]][-0-9:.+]*\)*$$||' >${PATCHDIR}_/patch-go.mod
	cd ${WRKSRC} && ${SETENV} TZ=UTC ${DIFF} -u --minimal -- go.sum.orig go.sum | \
		${SED} -e '/^---/s|\.[0-9]* +0000$$| UTC|' -e '/^+++/s|\([[:blank:]][-0-9:.+]*\)*$$||' >${PATCHDIR}_/patch-go.sum

post-deps-all: .PHONY
	${TEST} ! -d '${PATCHDIR}_' || ${MV} ${PATCHDIR}_ ${PATCHDIR}

pre-deps-all: .PHONY
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}_

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>

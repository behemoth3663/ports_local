PORTNAME=	pint
PORTVERSION=	0.88.0
DISTVERSIONPREFIX=	v
#PORTREVISION=	1
CATEGORIES=	local sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	Prometheus rule linter/validator
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	APACHE20

USES=		go:modules
GO_TARGET=	./cmd/${PORTNAME}

USE_GITHUB=	yes
GH_ACCOUNT=	cloudflare

USE_GITLAB=	nodefault

PLIST_FILES=	bin/${PORTNAME}

COMMIT_CMD=	git ls-remote \
		https://github.com/${GH_ACCOUNT}/${GH_PROJECT} \
		refs/tags/${DISTVERSIONFULL} | cut -c1-7

CGO_ENABLED=	0
GO_BUILDFLAGS+=	-ldflags="-s -w \
	-X main.version=${PORTVERSION} \
	-X main.commit=$$(${COMMIT_CMD})"
GO_ENV+=	GOEXPERIMENT=norandomizedheapbase64

pre-deps: .PHONY
	${TEST} -d ${PATCHDIR}_ && ${RM} -r ${PATCHDIR} || ${MV} ${PATCHDIR} ${PATCHDIR}_

post-deps-all: .PHONY
	${TEST} -d ${PATCHDIR}_ && ${MV} ${PATCHDIR}_ ${PATCHDIR} || :

pre-deps-all: .PHONY
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}_

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>

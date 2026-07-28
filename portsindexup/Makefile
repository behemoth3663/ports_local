PORTNAME=	portsindexup
PORTVERSION=	0.1.5
DISTVERSIONPREFIX=	v
#PORTREVISION=	1
CATEGORIES=	local ports-mgmt

MAINTAINER=	none@freebsd.org
COMMENT=	Update the FreeBSD ports INDEX file partially
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}/tree/main/cmd/${PORTNAME}

LICENSE=	BSD3CLAUSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	omilevskyi
GH_PROJECT=	go

WRKSRC_SUBDIR=	cmd/${PORTNAME}

PLIST_FILES=	bin/${PORTNAME}

COMMIT_CMD=	git ls-remote \
		https://github.com/${GH_ACCOUNT}/${GH_PROJECT} \
		refs/tags/${DISTVERSIONFULL} | cut -c1-7

CGO_ENABLED=	0
GO_BUILDFLAGS+=	-ldflags="-s -w \
		-X main.version=${PORTVERSION} \
		-X main.gitCommit=$$(${COMMIT_CMD}) \
		-X main.makeBin=/usr/bin/make"
GO_ENV+=	GOEXPERIMENT=norandomizedheapbase64

M2T_OFFLINE=	true
.export	M2T_OFFLINE

post-extract: .PHONY
	@${RM} ${WRKSRC}/../../go.work
	@${MKDIR} ${WRKSRC}/vendor
	@${CP} ${PATCHDIR}/modules.txt ${WRKSRC}/vendor

pre-deps: .PHONY
	${TEST} -d ${PATCHDIR}_ && ${RM} -r ${PATCHDIR} || ${MV} ${PATCHDIR} ${PATCHDIR}_

post-deps-all: .PHONY
	${TEST} ! -d ${PATCHDIR}_ || ${MV} ${PATCHDIR}_ ${PATCHDIR}

pre-deps-all: .PHONY
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}_

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>

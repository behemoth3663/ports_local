PORTNAME=	infracost
PORTVERSION=	0.10.30
DISTVERSIONPREFIX=	v
#PORTREVISION=	1
CATEGORIES=	local sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	Cloud cost estimates for Terraform
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	APACHE20
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules
GO_TARGET=	./cmd/${PORTNAME}

USE_GITHUB=	yes
GH_ACCOUNT=	infracost

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags '-static'"

PLIST_FILES=	bin/${PORTNAME}

M2T_OFFLINE=	true
.export	M2T_OFFLINE

_OPTIONS_extract=	600:pre-post-extract
pre-post-extract: .PHONY
	${FIND} ${WRKDIR} -maxdepth 1 -type d -name 'go-commons-*' | while read src; do \
		${LN} -s "$${src}" "$${src%%go-commons-*}gruntwork-cli-$${src##*go-commons-}"; \
	done

pre-deps: .PHONY .SILENT
	FILES=$$(${FIND} ${PATCHDIR} -type f -name 'patch-*'); \
	if [ -n "$${FILES}" ]; then \
		${TAR} -C / -cf ${.CURDIR}/${PORTNAME}.tar $${FILES}; \
		${RM} $${FILES}; \
	fi

pre-deps-all: .PHONY
post-deps-all: .PHONY
	${TAR} -C / -xf ${.CURDIR}/${PORTNAME}.tar
	${RM} ${PATCH_COOKIE} ${.CURDIR}/${PORTNAME}.tar

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>

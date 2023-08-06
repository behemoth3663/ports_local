PORTNAME=	infracost
PORTVERSION=	0.10.27
DISTVERSIONPREFIX=	v
#ORTREVISION=	1
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

_OPTIONS_extract=	600:pre-post-extract
pre-post-extract: .PHONY
	${FIND} ${WRKDIR} -maxdepth 1 -type d -name 'go-commons-*' | while read src; do \
		${LN} -s "$${src}" "$${src%%go-commons-*}gruntwork-cli-$${src##*go-commons-}"; \
	done

post-post-extract: .PHONY

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.-include "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>

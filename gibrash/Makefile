PORTNAME=	gibrash
PORTVERSION=	0.0.2
PORTREVISION=	1
DISTVERSIONPREFIX=	v
CATEGORIES=	local

MAINTAINER=	none@freebsd.org
COMMENT=	Print current git branch name for shell prompt
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	BSD3CLAUSE

USES=		llvm:22,build

BUILD_DEPENDS=	git:devel/git

USE_GITHUB=	yes
GH_ACCOUNT=	omilevskyi

# GH_PROJECT below is not expanded, so PORTNAME is used instead
GIT_COMMIT=	git ls-remote \
		https://github.com/${GH_ACCOUNT}/${PORTNAME} \
		refs/tags/${DISTVERSIONPREFIX}${PORTVERSION}

CFLAGS+=	-DVERSION='\"${PORTVERSION}\"' -DCOMMIT_HASH='\"${GIT_COMMIT:sh:C/(.{7}).*/\\1/W}\"'

PLIST_FILES=	bin/${PORTNAME}

post-patch: .PHONY
	${REINPLACE_CMD} \
		-e '/-fstack-protector/d' \
		-e '/-ftrivial-auto-var-init/d' \
		-e '/-fzero-call-used-regs/d' \
	${WRKSRC}/Makefile

.include <bsd.port.mk>

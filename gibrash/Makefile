PORTNAME=	gibrash
PORTVERSION=	0.0.2
PORTREVISION=	3
DISTVERSIONPREFIX=	v
CATEGORIES=	local

MAINTAINER=	none@freebsd.org
COMMENT=	Print current git branch name for shell prompt
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	BSD3CLAUSE

USES=		llvm:23,build

BUILD_DEPENDS=	git:devel/git

USE_GITHUB=	yes
GH_ACCOUNT=	omilevskyi

# GH_PROJECT is not expanded, use PORTNAME instead
COMMIT_CMD=	git ls-remote \
		https://github.com/${GH_ACCOUNT}/${PORTNAME} \
		refs/tags/${DISTVERSIONFULL}

CFLAGS+=       -DVERSION='\"${PORTVERSION}\"' -DCOMMIT_HASH='\"${COMMIT_CMD:sh:C/(.{7}).*/\\1/W}\"'

PLIST_FILES=	bin/${PORTNAME}

post-patch: .PHONY
	${REINPLACE_CMD} \
		-e '/-fstack-protector/d' \
		-e '/-ftrivial-auto-var-init/d' \
		-e '/-fzero-call-used-regs/d' \
	${WRKSRC}/Makefile

.include <bsd.port.mk>

PORTNAME=	Class-Refresh
PORTVERSION=	0.07
CATEGORIES=	devel perl5 local
MASTER_SITES=	CPAN
PKGNAMEPREFIX=	p5-

MAINTAINER=	none@FreeBSD.org
COMMENT=	Refresh your classes during runtime
WWW=		https://metacpan.org/release/Class-Refresh

LICENSE=	ART10 GPLv1+
LICENSE_COMB=	dual

RUN_DEPENDS=	\
	p5-Class-Load>=0:devel/p5-Class-Load \
	p5-Class-Unload>=0:devel/p5-Class-Unload \
	p5-Devel-OverrideGlobalRequire>=0:local/p5-Devel-OverrideGlobalRequire \
	p5-Try-Tiny>=0:lang/p5-Try-Tiny
TEST_DEPENDS=	\
	p5-Test-Fatal>=0:devel/p5-Test-Fatal \
	p5-Test-Requires>=0:devel/p5-Test-Requires

USES=		perl5
USE_PERL5=	configure

NO_ARCH=	yes

.include "../Mk/${USES:Mperl*}.mk"
.include <bsd.port.mk>

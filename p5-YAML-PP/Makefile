PORTNAME=	YAML-PP
PORTVERSION=	0.38.1
DISTVERSIONPREFIX=	v
CATEGORIES=	textproc perl5 local
MASTER_SITES=	CPAN
PKGNAMEPREFIX=	p5-

MAINTAINER=	dvl@FreeBSD.org
COMMENT=	YAML which aims to parse YAML 1.2 & conform to the spec
WWW=		https://github.com/perlpunk/YAML-PP-p5

LICENSE=	ART10 GPLv1+
LICENSE_COMB=	dual

TEST_DEPENDS=	p5-Test-Warn>=0:devel/p5-Test-Warn

USES=		perl5
USE_PERL5=	configure

NO_ARCH=	yes

.include "../Mk/${USES:Mperl*}.mk"
.include <bsd.port.mk>

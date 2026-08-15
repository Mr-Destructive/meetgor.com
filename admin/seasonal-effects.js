// Seasonal Effects - Auto-applied based on current date
// For testing: open console and run: localStorage.setItem('testMonth', 3); location.reload()
// Or use URL: ?testMonth=3
(function() {
    const today = new Date();
    let month = today.getMonth() + 1;
    let day = today.getDate();
    
    // Check for test month in localStorage or URL params
    const urlParams = new URLSearchParams(window.location.search);
    const testMonth = localStorage.getItem('testMonth') || urlParams.get('testMonth');
    if (testMonth) {
        month = parseInt(testMonth);
        day = 15; // Use middle of month for testing
        console.log('Seasonal Effects: TESTING MODE - Month:', month, 'Day:', day);
    }
    
    console.log('Seasonal Effects: Checking date - Month:', month, 'Day:', day);

    const isHome = document.querySelector('main.intro-section') !== null;
    if (isHome) {
        console.log('Seasonal Effects: Skipping effects on homepage');
        return;
    }
    
    // Track visibility state to prevent particle overflow when tab is hidden
    let isTabVisible = true;
    let animationId = null;
    let burstIntervalId = null;

    const articleEl = document.querySelector('article.blog-post, article.page-content');
    function isClickOutsideArticle(event) {
        if (!articleEl) return true;
        return !articleEl.contains(event.target);
    }

    function armActivation(startFn) {
        let activated = false;
        const handler = (event) => {
            if (activated) return;
            if (!isClickOutsideArticle(event)) return;
            activated = true;
            document.removeEventListener('click', handler);
            startFn();
        };
        document.addEventListener('click', handler, { passive: true });
    }
    
    document.addEventListener('visibilitychange', () => {
        isTabVisible = !document.hidden;
        if (!isTabVisible) {
            console.log('Seasonal Effects: Tab hidden, pausing animation');
        } else {
            console.log('Seasonal Effects: Tab visible, resuming animation');
        }
    });

    function isInRange(m, d, startMonth, startDay, endMonth, endDay) {
        if (startMonth === endMonth) {
            return m === startMonth && d >= startDay && d <= endDay;
        }
        if (m === startMonth) return d >= startDay;
        if (m === endMonth) return d <= endDay;
        return m > startMonth && m < endMonth;
    }

    function removeExistingEffects() {
        const canvas = document.getElementById('seasonal-canvas');
        if (canvas) canvas.remove();
    }

    function createCanvas() {
        const canvas = document.createElement('canvas');
        canvas.id = 'seasonal-canvas';
        canvas.style.position = 'fixed';
        canvas.style.top = '0';
        canvas.style.left = '0';
        canvas.style.zIndex = '0';
        canvas.style.width = '100%';
        canvas.style.height = '100%';
        canvas.style.pointerEvents = 'none';
        canvas.style.backgroundColor = 'transparent';
        document.body.appendChild(canvas);
        return canvas;
    }

    // December (Snowfall) - Dec 1 to Jan 5
    if ((month === 12) || (month === 1 && day <= 5)) {
        armActivation(() => {
            console.log('Seasonal Effects: Activating snowfall');
            removeExistingEffects();
            const canvas = createCanvas();
            const ctx = canvas.getContext('2d');
            const snowflakes = [];

            canvas.width = window.innerWidth;
            canvas.height = window.innerHeight;
            
            let animateSnow = null;

            class Snowflake {
                constructor() {
                    this.x = Math.random() * canvas.width;
                    this.y = Math.random() * canvas.height - canvas.height;
                    this.vx = Math.random() * 0.4 - 0.2;
                    this.vy = Math.random() * 0.6 + 0.4;
                    this.radius = Math.random() * 1.5 + 0.8;
                    this.opacity = Math.random() * 0.4 + 0.25;
                }

                update() {
                    this.x += this.vx;
                    this.y += this.vy;
                    if (this.y > canvas.height) {
                        this.y = -10;
                        this.x = Math.random() * canvas.width;
                    }
                }

                draw() {
                    ctx.fillStyle = 'rgba(200, 220, 255, ' + this.opacity + ')';
                    ctx.beginPath();
                    ctx.arc(this.x, this.y, this.radius, 0, Math.PI * 2);
                    ctx.fill();
                    ctx.strokeStyle = 'rgba(150, 200, 255, ' + (this.opacity * 0.6) + ')';
                    ctx.lineWidth = 0.5;
                    ctx.stroke();
                }
            }

            for (let i = 0; i < 18; i++) {
                snowflakes.push(new Snowflake());
            }

            function animate() {
                if (!isTabVisible) {
                    animateSnow = requestAnimationFrame(animate);
                    return;
                }
                ctx.clearRect(0, 0, canvas.width, canvas.height);
                for (let flake of snowflakes) {
                    flake.update();
                    flake.draw();
                }
                animateSnow = requestAnimationFrame(animate);
            }

            window.addEventListener('resize', () => {
                canvas.width = window.innerWidth;
                canvas.height = window.innerHeight;
            });

            animate();
        });
    }
    
    // Jan 10-20 (Makar Sankrant - Kites)
    else if (month === 1 && day >= 10 && day <= 20) {
        armActivation(() => {
            removeExistingEffects();
            const canvas = createCanvas();
            const ctx = canvas.getContext('2d');
            const kites = [];

            canvas.width = window.innerWidth;
            canvas.height = window.innerHeight;

            class Kite {
                constructor() {
                    this.x = Math.random() * canvas.width;
                    this.y = Math.random() * (canvas.height * 0.5);
                    this.vx = Math.random() * 0.6 - 0.3;
                    this.vy = Math.random() * 0.4 - 0.2;
                    this.size = Math.random() * 6 + 6;
                    this.rotation = Math.random() * Math.PI * 2;
                    this.rotSpeed = Math.random() * 0.008 - 0.004;
                    this.colors = ['#FF6B6B', '#FFD93D', '#6BCB77', '#4D96FF', '#FF6B9D'];
                    this.color = this.colors[Math.floor(Math.random() * this.colors.length)];
                }

                update() {
                    this.x += this.vx;
                    this.y += this.vy;
                    this.rotation += this.rotSpeed;
                    if (this.x > canvas.width) this.x = -30;
                    if (this.x < -30) this.x = canvas.width;
                    if (this.y > canvas.height) this.y = -30;
                }

                draw() {
                    ctx.save();
                    ctx.translate(this.x, this.y);
                    ctx.rotate(this.rotation);
                    
                    // Draw kite shape
                    ctx.fillStyle = this.color;
                    ctx.beginPath();
                    ctx.moveTo(0, -this.size);
                    ctx.lineTo(this.size * 0.6, 0);
                    ctx.lineTo(0, this.size * 0.8);
                    ctx.lineTo(-this.size * 0.6, 0);
                    ctx.closePath();
                    ctx.fill();
                    
                    // Draw kite border
                    ctx.strokeStyle = 'rgba(0,0,0,0.2)';
                    ctx.lineWidth = 0.5;
                    ctx.stroke();
                    
                    // Draw kite string
                    ctx.strokeStyle = 'rgba(100,100,100,0.3)';
                    ctx.lineWidth = 0.5;
                    ctx.beginPath();
                    ctx.moveTo(0, this.size * 0.8);
                    ctx.lineTo(this.size * 0.5, this.size * 1.5);
                    ctx.stroke();
                    
                    ctx.restore();
                }
            }

            for (let i = 0; i < 5; i++) {
                kites.push(new Kite());
            }

            function animate() {
                if (!isTabVisible) {
                    requestAnimationFrame(animate);
                    return;
                }
                ctx.clearRect(0, 0, canvas.width, canvas.height);
                for (let kite of kites) {
                    kite.update();
                    kite.draw();
                }
                requestAnimationFrame(animate);
            }

            window.addEventListener('resize', () => {
                canvas.width = window.innerWidth;
                canvas.height = window.innerHeight;
            });

            animate();
        });
    }
    
    // Mar 5-20 (Holi - Color Splash - Vibrant Interactive)
    else if (month === 3 && day >= 5 && day <= 20) {
        armActivation(() => {
            removeExistingEffects();
            const canvas = createCanvas();
            const ctx = canvas.getContext('2d');
            const particles = [];

            canvas.width = window.innerWidth;
            canvas.height = window.innerHeight;

            class HoliParticle {
                constructor(x, y) {
                    this.x = x || (Math.random() * canvas.width);
                    this.y = y || (Math.random() * canvas.height);
                    this.vx = Math.random() * 1.2 - 0.6;
                    this.vy = Math.random() * 1.2 - 0.6;
                    this.radius = Math.random() * 0.9 + 0.3;
                    this.opacity = 0.12;
                    this.colors = ['#FF1493', '#FFD700', '#00CED1', '#32CD32', '#FF6347', '#00FF7F', '#FF69B4', '#FF4500'];
                    this.color = this.colors[Math.floor(Math.random() * this.colors.length)];
                    this.life = 1;
                }

                update() {
                    this.x += this.vx;
                    this.y += this.vy;
                    this.vy += 0.05;
                    this.vx *= 0.99;
                    this.life -= 0.0035;
                    if (this.life < 0) this.life = 0;
                }

                draw() {
                    ctx.fillStyle = this.color.replace(')', ', ' + (this.life * this.opacity) + ')').replace('rgb(', 'rgba(');
                    ctx.beginPath();
                    ctx.arc(this.x, this.y, this.radius * Math.max(0.5, this.life), 0, Math.PI * 2);
                    ctx.fill();
                    
                    // Glow effect (reduced)
                    ctx.shadowColor = this.color;
                    ctx.shadowBlur = 4;
                    ctx.fillStyle = this.color.replace(')', ', ' + (this.life * this.opacity * 0.25) + ')').replace('rgb(', 'rgba(');
                    ctx.fill();
                    ctx.shadowBlur = 0;
                }
            }

            // Auto splash every 9 seconds (only if tab visible)
            burstIntervalId = setInterval(() => {
                if (!isTabVisible || particles.length > 70) return; // Skip if hidden or too many particles
                const splashX = Math.random() * canvas.width;
                const splashY = Math.random() * canvas.height * 0.6;
                for (let i = 0; i < 4; i++) {
                    particles.push(new HoliParticle(splashX, splashY));
                }
            }, 9000);

            document.addEventListener('mousemove', (e) => {
                if (!isClickOutsideArticle(e)) return;
                if (Math.random() > 0.95) {
                    let p = new HoliParticle(e.clientX, e.clientY);
                    particles.push(p);
                }
            });

            document.addEventListener('click', (e) => {
                if (!isTabVisible) return;
                if (!isClickOutsideArticle(e)) return;
                if (particles.length > 70) return; // Don't add if already too many
                for (let i = 0; i < 6; i++) {
                    let p = new HoliParticle(e.clientX, e.clientY);
                    particles.push(p);
                }
            });

            function animate() {
                if (!isTabVisible) {
                    animationId = requestAnimationFrame(animate);
                    return;
                }
                
                ctx.clearRect(0, 0, canvas.width, canvas.height);
                
                // Limit particles to prevent overflow
                if (particles.length > 200) {
                    particles.splice(0, particles.length - 200);
                }
                
                for (let i = particles.length - 1; i >= 0; i--) {
                    particles[i].update();
                    if (particles[i].life <= 0) {
                        particles.splice(i, 1);
                    } else {
                        particles[i].draw();
                    }
                }
                animationId = requestAnimationFrame(animate);
            }

            window.addEventListener('resize', () => {
                canvas.width = window.innerWidth;
                canvas.height = window.innerHeight;
            });

            animate();
        });
    }
    
    // April-May (Mango Blossom - Soft Petals)
    else if (month === 4 || month === 5) {
        armActivation(() => {
            removeExistingEffects();
            const canvas = createCanvas();
            const ctx = canvas.getContext('2d');
            const petals = [];

            canvas.width = window.innerWidth;
            canvas.height = window.innerHeight;

            class Petal {
                constructor() {
                    this.x = Math.random() * canvas.width;
                    this.y = Math.random() * canvas.height - canvas.height;
                    this.vx = Math.random() * 0.35 - 0.18;
                    this.vy = Math.random() * 0.6 + 0.3;
                    this.rotation = Math.random() * Math.PI * 2;
                    this.rotationSpeed = Math.random() * 0.06 - 0.03;
                    this.size = Math.random() * 4 + 3;
                    this.colors = ['#FFF6D5', '#FFE8A3', '#FCE7B1', '#FFF0C9', '#FDEAC2'];
                    this.color = this.colors[Math.floor(Math.random() * this.colors.length)];
                }

                update() {
                    this.x += this.vx;
                    this.y += this.vy;
                    this.rotation += this.rotationSpeed;
                    this.vx += Math.random() * 0.06 - 0.03;
                    if (this.vx > 0.4) this.vx = 0.4;
                    if (this.vx < -0.4) this.vx = -0.4;
                    if (this.y > canvas.height) {
                        this.y = -10;
                        this.x = Math.random() * canvas.width;
                    }
                }

                draw() {
                    ctx.save();
                    ctx.translate(this.x, this.y);
                    ctx.rotate(this.rotation);
                    
                    // Draw petal shape
                    ctx.fillStyle = this.color;
                    ctx.beginPath();
                    ctx.ellipse(0, 0, this.size, this.size * 0.6, 0, 0, Math.PI * 2);
                    ctx.fill();
                    
                    ctx.restore();
                }
            }

            for (let i = 0; i < 14; i++) {
                petals.push(new Petal());
            }

            function animate() {
                if (!isTabVisible) {
                    requestAnimationFrame(animate);
                    return;
                }
                ctx.clearRect(0, 0, canvas.width, canvas.height);
                for (let petal of petals) {
                    petal.update();
                    petal.draw();
                }
                requestAnimationFrame(animate);
            }

            window.addEventListener('resize', () => {
                canvas.width = window.innerWidth;
                canvas.height = window.innerHeight;
            });

            animate();
        });
    }
    
    // June (Summer Heat - Warm Glimmer)
    else if (month === 6) {
        armActivation(() => {
            removeExistingEffects();
            const canvas = createCanvas();
            const ctx = canvas.getContext('2d');
            const glimmers = [];

            canvas.width = window.innerWidth;
            canvas.height = window.innerHeight;

            class Glimmer {
                constructor() {
                    this.x = Math.random() * canvas.width;
                    this.y = Math.random() * canvas.height * 0.7;
                    this.vx = Math.random() * 0.2 - 0.1;
                    this.vy = Math.random() * 0.2 - 0.1;
                    this.radius = Math.random() * 1.2 + 0.6;
                    this.opacity = Math.random() * 0.08 + 0.04;
                    this.colors = ['#FFD9A0', '#FFE2B8', '#FFC987', '#FFD1A8'];
                    this.color = this.colors[Math.floor(Math.random() * this.colors.length)];
                }

                update() {
                    this.x += this.vx;
                    this.y += this.vy;
                    if (this.x < -10) this.x = canvas.width + 10;
                    if (this.x > canvas.width + 10) this.x = -10;
                    if (this.y < -10) this.y = canvas.height * 0.7 + 10;
                    if (this.y > canvas.height * 0.7 + 10) this.y = -10;
                }

                draw() {
                    ctx.fillStyle = this.color.replace(')', ', ' + this.opacity + ')').replace('rgb(', 'rgba(');
                    ctx.beginPath();
                    ctx.arc(this.x, this.y, this.radius, 0, Math.PI * 2);
                    ctx.fill();
                }
            }

            for (let i = 0; i < 12; i++) {
                glimmers.push(new Glimmer());
            }

            function animate() {
                if (!isTabVisible) {
                    requestAnimationFrame(animate);
                    return;
                }
                ctx.clearRect(0, 0, canvas.width, canvas.height);
                for (let g of glimmers) {
                    g.update();
                    g.draw();
                }
                requestAnimationFrame(animate);
            }

            window.addEventListener('resize', () => {
                canvas.width = window.innerWidth;
                canvas.height = window.innerHeight;
            });

            animate();
        });
    }

    // July-August (Monsoon - Light Rain)
    else if ((month >= 7 && month <= 8)) {
        armActivation(() => {
            removeExistingEffects();
            const canvas = createCanvas();
            const ctx = canvas.getContext('2d');
            const raindrops = [];

            canvas.width = window.innerWidth;
            canvas.height = window.innerHeight;

            class Raindrop {
                constructor() {
                    this.x = Math.random() * canvas.width;
                    this.y = Math.random() * canvas.height - canvas.height;
                    this.vy = Math.random() * 3 + 3.5;
                    this.vx = Math.random() * 0.6 - 0.3;
                    this.length = Math.random() * 10 + 8;
                    this.opacity = Math.random() * 0.4 + 0.2;
                }

                update() {
                    this.x += this.vx;
                    this.y += this.vy;
                    if (this.y > canvas.height) {
                        this.y = -10;
                        this.x = Math.random() * canvas.width;
                    }
                }

                draw() {
                    ctx.strokeStyle = 'rgba(100, 150, 200, ' + this.opacity + ')';
                    ctx.lineWidth = 1;
                    ctx.beginPath();
                    ctx.moveTo(this.x, this.y);
                    ctx.lineTo(this.x + this.vx, this.y + this.length);
                    ctx.stroke();
                }
            }

            for (let i = 0; i < 45; i++) {
                raindrops.push(new Raindrop());
            }

            function animate() {
                if (!isTabVisible) {
                    requestAnimationFrame(animate);
                    return;
                }
                ctx.clearRect(0, 0, canvas.width, canvas.height);
                for (let drop of raindrops) {
                    drop.update();
                    drop.draw();
                }
                requestAnimationFrame(animate);
            }

            window.addEventListener('resize', () => {
                canvas.width = window.innerWidth;
                canvas.height = window.innerHeight;
            });

            animate();
        });
    }
    
    // September (Ganesh Chaturthi / Onam - Floating Marigolds)
    else if (month === 9) {
        armActivation(() => {
            removeExistingEffects();
            const canvas = createCanvas();
            const ctx = canvas.getContext('2d');
            const flowers = [];

            canvas.width = window.innerWidth;
            canvas.height = window.innerHeight;

            class Flower {
                constructor() {
                    this.x = Math.random() * canvas.width;
                    this.y = Math.random() * canvas.height - canvas.height;
                    this.vx = Math.random() * 0.3 - 0.15;
                    this.vy = Math.random() * 0.6 + 0.2;
                    this.radius = Math.random() * 2.5 + 1.5;
                    this.opacity = Math.random() * 0.25 + 0.15;
                    this.colors = ['#FFB000', '#FFA000', '#FF8C42', '#FFD166'];
                    this.color = this.colors[Math.floor(Math.random() * this.colors.length)];
                }

                update() {
                    this.x += this.vx;
                    this.y += this.vy;
                    if (this.y > canvas.height) {
                        this.y = -10;
                        this.x = Math.random() * canvas.width;
                    }
                }

                draw() {
                    ctx.fillStyle = this.color.replace(')', ', ' + this.opacity + ')').replace('rgb(', 'rgba(');
                    ctx.beginPath();
                    ctx.arc(this.x, this.y, this.radius, 0, Math.PI * 2);
                    ctx.fill();
                }
            }

            for (let i = 0; i < 18; i++) {
                flowers.push(new Flower());
            }

            function animate() {
                if (!isTabVisible) {
                    requestAnimationFrame(animate);
                    return;
                }
                ctx.clearRect(0, 0, canvas.width, canvas.height);
                for (let f of flowers) {
                    f.update();
                    f.draw();
                }
                requestAnimationFrame(animate);
            }

            window.addEventListener('resize', () => {
                canvas.width = window.innerWidth;
                canvas.height = window.innerHeight;
            });

            animate();
        });
    }
    
    
    // October-November (Navratri/Diwali - Gentle Firecrackers)
    else if (month === 10 || month === 11) {
        armActivation(() => {
            removeExistingEffects();
            const canvas = createCanvas();
            const ctx = canvas.getContext('2d');
            const particles = [];

            canvas.width = window.innerWidth;
            canvas.height = window.innerHeight;

            class DiwaliFirecracker {
                constructor(x, y) {
                    this.x = x || (Math.random() * canvas.width);
                    this.y = y || (Math.random() * (canvas.height * 0.5) + canvas.height * 0.1);
                    this.vx = Math.random() * 1.8 - 0.9;
                    this.vy = Math.random() * 1.8 - 0.9;
                    this.radius = Math.random() * 0.9 + 0.25;
                    this.opacity = 0.1;
                    this.life = 1;
                    this.colors = ['#FF0000', '#FFD700', '#FF6B00', '#FF1493', '#FF4500', '#FFFF00', '#FF8C00', '#00FFFF'];
                    this.color = this.colors[Math.floor(Math.random() * this.colors.length)];
                }

                update() {
                    this.x += this.vx;
                    this.y += this.vy;
                    this.vy += 0.08;
                    this.vx *= 0.98;
                    this.life -= 0.007;
                    if (this.life < 0) this.life = 0;
                }

                draw() {
                    // Subtle core
                    ctx.fillStyle = this.color.replace(')', ', ' + (this.life * this.opacity) + ')').replace(/rgb\(/g, 'rgba(');
                    ctx.beginPath();
                    ctx.arc(this.x, this.y, this.radius * Math.max(0.3, this.life), 0, Math.PI * 2);
                    ctx.fill();
                    
                    // Light glow (reduced)
                    ctx.shadowColor = this.color;
                    ctx.shadowBlur = 3;
                    ctx.fillStyle = this.color.replace(')', ', ' + (this.life * this.opacity * 0.18) + ')').replace(/rgb\(/g, 'rgba(');
                    ctx.fill();
                    ctx.shadowBlur = 0;
                }
            }

            // Burst firecrackers less frequently (only if tab visible)
            burstIntervalId = setInterval(() => {
                if (!isTabVisible || particles.length > 50) return; // Skip if hidden or too many particles
                const burstX = Math.random() * canvas.width;
                const burstY = Math.random() * (canvas.height * 0.4) + canvas.height * 0.15;
                for (let i = 0; i < 6; i++) {
                    particles.push(new DiwaliFirecracker(burstX, burstY));
                }
            }, 4500);

            function animate() {
                if (!isTabVisible) {
                    animationId = requestAnimationFrame(animate);
                    return;
                }
                
                ctx.clearRect(0, 0, canvas.width, canvas.height);
                
                // Limit particles to prevent overflow
                if (particles.length > 600) {
                    particles.splice(0, particles.length - 600);
                }
                
                for (let i = particles.length - 1; i >= 0; i--) {
                    particles[i].update();
                    if (particles[i].life <= 0) {
                        particles.splice(i, 1);
                    } else {
                        particles[i].draw();
                    }
                }
                animationId = requestAnimationFrame(animate);
            }

            window.addEventListener('resize', () => {
                canvas.width = window.innerWidth;
                canvas.height = window.innerHeight;
            });

            animate();
        });
    }
})();

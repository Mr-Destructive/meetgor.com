package plugins

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"

	"github.com/mr-destructive/mr-destructive.github.io/models"
)

type SeasonalEffectsPlugin struct {
	PluginName string
}

func (p *SeasonalEffectsPlugin) Name() string {
	return p.PluginName
}

func (p *SeasonalEffectsPlugin) Execute(ssg *models.SSG) {
	config := &ssg.Config
	outputPath := filepath.Join(".", config.Blog.OutputDir)
	
	// Create the effects JavaScript file
	effectsJS := generateSeasonalEffectsJS()
	
	effectsPath := filepath.Join(outputPath, "seasonal-effects.js")
	err := os.WriteFile(effectsPath, []byte(effectsJS), 0660)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println("Seasonal effects plugin: Generated seasonal-effects.js")
}

func generateSeasonalEffectsJS() string {
	return `// Seasonal Effects - Auto-applied based on current date
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
    
    // Track visibility state to prevent particle overflow when tab is hidden
    let isTabVisible = true;
    let animationId = null;
    let burstIntervalId = null;
    
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
        canvas.style.width = '100%';
        canvas.style.height = '100%';
        canvas.style.pointerEvents = 'none';
        canvas.style.zIndex = '-1';
        canvas.style.backgroundColor = 'transparent';
        document.body.appendChild(canvas);
        return canvas;
    }

    // December (Snowfall) - Dec 1 to Jan 5
    if ((month === 12) || (month === 1 && day <= 5)) {
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
                this.vx = Math.random() * 0.5 - 0.25;
                this.vy = Math.random() * 0.8 + 0.5;
                this.radius = Math.random() * 2 + 1;
                this.opacity = Math.random() * 0.5 + 0.3;
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

        for (let i = 0; i < 30; i++) {
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
    }
    
    // Jan 10-20 (Makar Sankrant - Kites)
    else if (month === 1 && day >= 10 && day <= 20) {
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
                this.vx = Math.random() * 0.8 - 0.4;
                this.vy = Math.random() * 0.5 - 0.25;
                this.size = Math.random() * 8 + 8;
                this.rotation = Math.random() * Math.PI * 2;
                this.rotSpeed = Math.random() * 0.01 - 0.005;
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

        for (let i = 0; i < 8; i++) {
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
    }
    
    // Mar 5-20 (Holi - Color Splash - Vibrant Interactive)
    else if (month === 3 && day >= 5 && day <= 20) {
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
                this.vx = Math.random() * 2 - 1;
                this.vy = Math.random() * 2 - 1;
                this.radius = Math.random() * 1.2 + 0.4;
                this.opacity = 0.15;
                this.colors = ['#FF1493', '#FFD700', '#00CED1', '#32CD32', '#FF6347', '#00FF7F', '#FF69B4', '#FF4500'];
                this.color = this.colors[Math.floor(Math.random() * this.colors.length)];
                this.life = 1;
            }

            update() {
                this.x += this.vx;
                this.y += this.vy;
                this.vy += 0.06;
                this.vx *= 0.99;
                this.life -= 0.004;
                if (this.life < 0) this.life = 0;
            }

            draw() {
                ctx.fillStyle = this.color.replace(')', ', ' + (this.life * this.opacity) + ')').replace('rgb(', 'rgba(');
                ctx.beginPath();
                ctx.arc(this.x, this.y, this.radius * Math.max(0.5, this.life), 0, Math.PI * 2);
                ctx.fill();
                
                // Glow effect (reduced)
                ctx.shadowColor = this.color;
                ctx.shadowBlur = 6;
                ctx.fillStyle = this.color.replace(')', ', ' + (this.life * this.opacity * 0.3) + ')').replace('rgb(', 'rgba(');
                ctx.fill();
                ctx.shadowBlur = 0;
            }
        }

        // Auto splash every 6 seconds (only if tab visible)
        burstIntervalId = setInterval(() => {
            if (!isTabVisible || particles.length > 100) return; // Skip if hidden or too many particles
            const splashX = Math.random() * canvas.width;
            const splashY = Math.random() * canvas.height * 0.6;
            for (let i = 0; i < 6; i++) {
                particles.push(new HoliParticle(splashX, splashY));
            }
        }, 6000);

        document.addEventListener('mousemove', (e) => {
            if (Math.random() > 0.92) {
                for (let i = 0; i < 2; i++) {
                    let p = new HoliParticle(e.clientX, e.clientY);
                    particles.push(p);
                }
            }
        });

        document.addEventListener('click', (e) => {
            if (!isTabVisible) return;
            if (particles.length > 100) return; // Don't add if already too many
            for (let i = 0; i < 10; i++) {
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
            if (particles.length > 300) {
                particles.splice(0, particles.length - 300);
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
    }
    
    // April (Spring - Falling Leaves)
    else if (month === 4) {
        removeExistingEffects();
        const canvas = createCanvas();
        const ctx = canvas.getContext('2d');
        const leaves = [];

        canvas.width = window.innerWidth;
        canvas.height = window.innerHeight;

        class Leaf {
            constructor() {
                this.x = Math.random() * canvas.width;
                this.y = Math.random() * canvas.height - canvas.height;
                this.vx = Math.random() * 0.5 - 0.25;
                this.vy = Math.random() * 0.8 + 0.5;
                this.rotation = Math.random() * Math.PI * 2;
                this.rotationSpeed = Math.random() * 0.08 - 0.04;
                this.size = Math.random() * 6 + 4;
                this.colors = ['#90EE90', '#98FB98', '#8FBC8F', '#3CB371', '#228B22', '#7CFC00'];
                this.color = this.colors[Math.floor(Math.random() * this.colors.length)];
            }

            update() {
                this.x += this.vx;
                this.y += this.vy;
                this.rotation += this.rotationSpeed;
                this.vx += Math.random() * 0.1 - 0.05;
                if (this.vx > 0.5) this.vx = 0.5;
                if (this.vx < -0.5) this.vx = -0.5;
                if (this.y > canvas.height) {
                    this.y = -10;
                    this.x = Math.random() * canvas.width;
                }
            }

            draw() {
                ctx.save();
                ctx.translate(this.x, this.y);
                ctx.rotate(this.rotation);
                
                // Draw leaf shape
                ctx.fillStyle = this.color;
                ctx.beginPath();
                ctx.ellipse(0, 0, this.size, this.size * 0.6, 0, 0, Math.PI * 2);
                ctx.fill();
                
                // Draw leaf vein
                ctx.strokeStyle = 'rgba(100, 150, 100, 0.3)';
                ctx.lineWidth = 0.5;
                ctx.beginPath();
                ctx.moveTo(0, -this.size * 0.5);
                ctx.lineTo(0, this.size * 0.5);
                ctx.stroke();
                
                ctx.restore();
            }
        }

        for (let i = 0; i < 20; i++) {
            leaves.push(new Leaf());
        }

        function animate() {
            if (!isTabVisible) {
                requestAnimationFrame(animate);
                return;
            }
            ctx.clearRect(0, 0, canvas.width, canvas.height);
            for (let leaf of leaves) {
                leaf.update();
                leaf.draw();
            }
            requestAnimationFrame(animate);
        }

        window.addEventListener('resize', () => {
            canvas.width = window.innerWidth;
            canvas.height = window.innerHeight;
        });

        animate();
    }
    
    // July-August (Monsoon - Light Rain)
    else if ((month >= 7 && month <= 8)) {
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
                this.vy = Math.random() * 4 + 5;
                this.vx = Math.random() * 1 - 0.5;
                this.length = Math.random() * 15 + 10;
                this.opacity = Math.random() * 0.5 + 0.3;
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
                ctx.lineWidth = 2;
                ctx.beginPath();
                ctx.moveTo(this.x, this.y);
                ctx.lineTo(this.x + this.vx, this.y + this.length);
                ctx.stroke();
            }
        }

        for (let i = 0; i < 80; i++) {
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
    }
    
    // October 25 - November 5 (Halloween - Spooky)
    else if ((month === 10 && day >= 25) || (month === 11 && day <= 5)) {
        removeExistingEffects();
        const canvas = createCanvas();
        const ctx = canvas.getContext('2d');
        const pumpkins = [];
        const bats = [];

        canvas.width = window.innerWidth;
        canvas.height = window.innerHeight;

        class Pumpkin {
            constructor() {
                this.x = Math.random() * canvas.width;
                this.y = Math.random() * canvas.height;
                this.size = Math.random() * 8 + 6;
                this.opacity = Math.random() * 0.15 + 0.08;
            }

            draw() {
                ctx.save();
                ctx.fillStyle = 'rgba(255, 140, 0, ' + this.opacity + ')';
                
                // Draw pumpkin body (4 bumps)
                for (let i = 0; i < 4; i++) {
                    const angle = (i / 4) * Math.PI * 2;
                    const bumpX = this.x + Math.cos(angle) * this.size * 0.6;
                    const bumpY = this.y + Math.sin(angle) * this.size * 0.6;
                    ctx.beginPath();
                    ctx.arc(bumpX, bumpY, this.size * 0.5, 0, Math.PI * 2);
                    ctx.fill();
                }
                
                // Draw eyes
                ctx.fillStyle = 'rgba(0, 0, 0, ' + (this.opacity * 0.8) + ')';
                ctx.beginPath();
                ctx.arc(this.x - this.size * 0.3, this.y - this.size * 0.2, this.size * 0.15, 0, Math.PI * 2);
                ctx.fill();
                ctx.beginPath();
                ctx.arc(this.x + this.size * 0.3, this.y - this.size * 0.2, this.size * 0.15, 0, Math.PI * 2);
                ctx.fill();
                
                ctx.restore();
            }
        }

        class Bat {
            constructor() {
                this.x = Math.random() * canvas.width;
                this.y = Math.random() * canvas.height * 0.4;
                this.vx = Math.random() * 0.8 - 0.4;
                this.vy = Math.random() * 0.3 - 0.15;
                this.size = Math.random() * 4 + 3;
                this.opacity = Math.random() * 0.2 + 0.1;
            }

            update() {
                this.x += this.vx;
                this.y += this.vy;
                if (this.x < -20) this.x = canvas.width + 20;
                if (this.x > canvas.width + 20) this.x = -20;
            }

            draw() {
                ctx.save();
                ctx.fillStyle = 'rgba(30, 30, 30, ' + this.opacity + ')';
                ctx.translate(this.x, this.y);
                
                // Draw bat wings
                ctx.beginPath();
                ctx.moveTo(0, 0);
                ctx.lineTo(-this.size * 0.8, -this.size * 0.5);
                ctx.lineTo(-this.size * 0.6, this.size * 0.3);
                ctx.closePath();
                ctx.fill();
                
                ctx.beginPath();
                ctx.moveTo(0, 0);
                ctx.lineTo(this.size * 0.8, -this.size * 0.5);
                ctx.lineTo(this.size * 0.6, this.size * 0.3);
                ctx.closePath();
                ctx.fill();
                
                ctx.restore();
            }
        }

        for (let i = 0; i < 5; i++) {
            pumpkins.push(new Pumpkin());
        }
        for (let i = 0; i < 4; i++) {
            bats.push(new Bat());
        }

        function animate() {
            if (!isTabVisible) {
                requestAnimationFrame(animate);
                return;
            }
            ctx.clearRect(0, 0, canvas.width, canvas.height);
            
            for (let pumpkin of pumpkins) {
                pumpkin.draw();
            }
            
            for (let bat of bats) {
                bat.update();
                bat.draw();
            }
            
            requestAnimationFrame(animate);
        }

        window.addEventListener('resize', () => {
            canvas.width = window.innerWidth;
            canvas.height = window.innerHeight;
        });

        animate();
    }
    
    // October (except 25+) and November 6+ (Diwali - Vibrant Firecrackers)
    else if ((month === 10 && day < 25) || (month === 11 && day > 5)) {
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
                this.vx = Math.random() * 2.5 - 1.25;
                this.vy = Math.random() * 2.5 - 1.25;
                this.radius = Math.random() * 1 + 0.3;
                this.opacity = 0.12;
                this.life = 1;
                this.colors = ['#FF0000', '#FFD700', '#FF6B00', '#FF1493', '#FF4500', '#FFFF00', '#FF8C00', '#00FFFF'];
                this.color = this.colors[Math.floor(Math.random() * this.colors.length)];
            }

            update() {
                this.x += this.vx;
                this.y += this.vy;
                this.vy += 0.1;
                this.vx *= 0.98;
                this.life -= 0.006;
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
                ctx.shadowBlur = 4;
                ctx.fillStyle = this.color.replace(')', ', ' + (this.life * this.opacity * 0.2) + ')').replace(/rgb\(/g, 'rgba(');
                ctx.fill();
                ctx.shadowBlur = 0;
            }
        }

        // Burst firecrackers less frequently (only if tab visible)
        burstIntervalId = setInterval(() => {
            if (!isTabVisible || particles.length > 80) return; // Skip if hidden or too many particles
            const burstX = Math.random() * canvas.width;
            const burstY = Math.random() * (canvas.height * 0.4) + canvas.height * 0.15;
            // Single burst instead of double
            for (let i = 0; i < 8; i++) {
                particles.push(new DiwaliFirecracker(burstX, burstY));
            }
        }, 3500);

        function animate() {
            if (!isTabVisible) {
                animationId = requestAnimationFrame(animate);
                return;
            }
            
            ctx.clearRect(0, 0, canvas.width, canvas.height);
            
            // Limit particles to prevent overflow
            if (particles.length > 1500) {
                particles.splice(0, particles.length - 1500);
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
    }
})();
`
}

func init() {
	RegisterPlugin("SeasonalEffects", reflect.TypeOf(SeasonalEffectsPlugin{}))
}

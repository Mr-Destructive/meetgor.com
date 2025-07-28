{
  "title": "Just Fucking use kubernetes",
  "post_dir": "thoughts",
  "type": "thoughts",
  "status": "published",
  "slug": "just-fucking-use-kubernetes",
  "date": "2025-07-28"
}

Just fucking use Kubernetes - [https://waylonwalker.com/just-fucking-use-kubernetes/](https://waylonwalker.com/just-fucking-use-kubernetes/)


This is the opposite post of [this](https://sliplane.io/blog/kubernetes-isnt-for-you), that I shared my thoughts on [here](https://www.meetgor.com/thoughts/kubernetes-isn-t-for-you/). 

I had mostly agreed to not use Kubernetes, of course I stated it depends. And this post just perfectly summarizes what those conditions of when to use kubernetes are. 

- You need to scale on day 1 (your app is small but not your ambitions) 
- Reliability and Rollback

Yes, those are very specific requirements to use kubernetes. But those are not just 2 needs, there are more and even I don't know when can you pull the trigger to shift to kubernetes. 

I want to emphasize on this point: 
> "But my app is small, so is your ambition"

That just hits hard, well done AI. This is me, when writing and building apps, I start with some things think as it was a overkill for that scale. 

But then it hits again with this quote: 
> What if it’s overkill? What if YOU are underkill

OK, this is the point that took a U-Turn at this post. I went from I think I disagree, to ok, I completely agree with this. 

Sometimes, all you need is a mindset shift, a blocker in your mind that holds you back from doing certain things. And for me, I have consumed enough tutorials and posts about Kubernetes, that I need to put to use and create. I have been stuck in the learning cycle, lets push to prod with kubernetes. 

I totally don't get the point of "Kubernetes is too heavy", there are slim and lightweight distributions like k3s, k0s, micro k8s, minikube, and what not. Those are not random things, they are built and maintained by large and reputed companies. 

If you think Kubernetes is too small for you, open shift, Rancher, GKE, AKS, EKS, and all other heavy weight distributions are available too, if 

Everyone is using it, kubernetes is widely adopted, true, but if you don't think its worth it, don't use it. But then still, if you have complexity and avoiding it, you'd be rewriting kubernetes yourself. 

I agree to this post too, correct for pointing out the low ambitions, complaining about problems that are already solved, laziness to explore curiosity, and missing a chance to be a nerd.

Does anyone care if you use simple yet fragile bash scripts or heavy weight Kubernetes cluster for just clicking buttons and creating and updating rows in a database? No! 

You know what, let's fucking use Kubernetes. 


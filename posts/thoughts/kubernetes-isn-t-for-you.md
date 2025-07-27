{
  "title": "Kubernetes isn't for you",
  "post_dir": "thoughts",
  "type": "thoughts",
  "status": "published",
  "slug": "kubernetes-isn-t-for-you",
  "date": "2025-07-27"
}

[https://sliplane.io/blog/kubernetes-isnt-for-you](https://sliplane.io/blog/kubernetes-isnt-for-you)

dev.to link: https://dev.to/code42cate/kubernetes-isnt-for-you-2c2m

## Kubernetes isn't for you

I almost agree with this, if your end goal is to get an app or product up and running as fast as possible.

The post rightly mentioned the points that come in the way of deployment: 

- Complexity and Learning curve
- Managing the platform instead of building the product

I think these two are good enough reasons to not use Kubernetes. 

I also agee with these 2 points strongly. I have seen and experienced it myself at my current company. It is a small product, not more than 1k customers, and using kubernetes, is that really needed? 

I think one of the 2 things might have caught their minds. 

- Big companies use kubernetes, we should too
- It feels professional, it sounds like we know what we are doing

People are caught with FOMO (fear of missing out) and also an egoistic approach in developing a product especially when they are small team, no oke is there to leash the technical decisions. 

Why can't we just keep it simple, just use docker swarm, just an image, if needed use multiple cloud functions. Why services? If all of them use the same backend in some way? 

Its not a good advice to not use kubernetes for all, I agree. If you are a large team, building a product with many services, across regios, maybe having different foundations or flows, Kubernetes makes sense. But for small teams, I do not think its worth it. 

I may be wrong here. But I am not 100% in agreement with it, it as usual as everything in tech, it depends. 

- For startups with limited services, maybe a monolithic structures, Kubernetes could be clearly avoided.
- For big startups, or large distributed companies, having many services and backends, Kubernetes makes sense.

If you want to use it for the purpose of learning it, please do use it. 

Kubernetes as usual is a tool like others, you can't use one tool everywhere. Where bash scripts work, they just work, where they don't they fall apart too, kubernetes works like a charm. 

Use your grug brains a little and choose wisely! In the end, who the hell cares if you use kubernetes or bash scripts to scale if your users are happy? 

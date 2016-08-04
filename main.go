package main

import (
    "fmt"
    "math"
    "os"

    "github.com/Raytracer/src"
)

func panicIf(err error) {
    if err != nil {
        panic(err)
    }
}

func hit_sphere(center utils.Vector, radius float64, r *utils.Ray) float64 {
    oc := utils.Subtract(r.Origin(), center)
    a := utils.Dot(r.Direction(), r.Direction())
    b := 2 * utils.Dot(oc, r.Direction())
    c := utils.Dot(oc, oc) - radius*radius
    dis := b*b - 4*a*c
    if dis < 0 {
        return -1
    } else {
        return (-b - math.Sqrt(dis)) / (2*a)
    }
}

func color(r *utils.Ray) utils.Vector {
    t := hit_sphere(utils.Vector{0, 0, -5}, 2.5, r)
    if t > 0{
        fmt.Println(r)
        N := utils.MakeUnitVector(utils.Subtract(r.PointAtParameter(t), utils.Vector{0, 0, -1}))
        return utils.MultiplyScalar(utils.Vector{N.X() + 1, N.Y() + 1, N.Z() + 1}, 0.5)
    }

    unit_direction := utils.MakeUnitVector(r.Direction())
    t = 0.5 * (unit_direction.Y() + 1)
    return utils.Add( utils.MultiplyScalar(utils.Vector{1, 1, 1}, 1 - t), utils.MultiplyScalar(utils.Vector{0.5, 0.7, 1}, t))
}

func main() {
    file, err := os.Create("image.ppm")
    panicIf(err)

    defer func() {
        err := file.Close()
        panicIf(err)
    }()

    nx, ny := 1000, 500
    fmt.Fprintf(file, "P3\n%d %d\n255\n", nx, ny);

    lowerLeft := utils.Vector{-10, -5, -5}
    horizontal := utils.Vector{20, 0, 0}
    vertical := utils.Vector{0, 10, 0}
    origin := utils.Vector{0, 0, 0}
    for j := ny-1; j >=0; j-- {
        for i := 0; i < nx; i++ {
            u := float64(i) / float64(nx)
            v := float64(j) / float64(ny)
            r := utils.Ray{origin, utils.Add(lowerLeft, utils.Add( utils.MultiplyScalar(horizontal, u),
                                                utils.MultiplyScalar(vertical, v)))}
            c := color(&r)
            ir := int(255.99*c.R())
            ig := int(255.99*c.G())
            ib := int(255.99*c.B())

            fmt.Fprintf(file, "%d %d %d\n", ir, ig, ib)
        }
    }
}

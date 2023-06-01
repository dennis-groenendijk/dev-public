<?php

class Subscription
{
    protected Gateway $gateway;

    public function __contruct(Gateway $gateway)
    {
        $this->gateway = $gateway;
    }

    public function create()
    {

    }

    public function cancel()
    {
        // find {payment service} customer
        //$customer = $this->gateway->findStripeCustomer();

        // find {payment service} subscription by customer
        //$this->gateway->findStripeSubscriptionByCustomer();

        $this->gateway->findCustomer();
    }

    public function invoice()
    {

    }

    public function swap($newPlan)
    {

    }
}

interface Gateway
{
    public function findCustomer();

    public function findSubscriptionByCustomer();
}

// Stripe is online billing service
class StripeGateway implements Gateway
{
    public function findCustomer()
    {

    }

    public function findSubscriptionByCustomer()
    {

    }
}

new Subscription(new StripeGateway());
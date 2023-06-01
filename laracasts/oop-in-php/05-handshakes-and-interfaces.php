<?php

interface Newsletter
{
    public function subscribe($email);
}

class CampaignMonitor implements Newsletter
{
//    protected $apiKey;
//
//    public function __construct($apiKey)
//    {
//        $this->apiKey = $apiKey;
//    }

    public function subscribe($email)
    {
        die('subscribing with CampaignMonitor');
    }
}

class Drip implements Newsletter
{
    public function subscribe($email)
    {
        die('subscribing with Drip');
    }
}

class NewsletterSubscriptionsController
{
    public function store(Newsletter $newsletter)
    {
        $email = 'joe@example.com';

        $newsletter->subscribe($email);
        //$newsletter->subscribe(auth()->user()->email); // Laravel auth
    }
}

$controller = new NewsletterSubscriptionsController();

$controller->store(new Drip());

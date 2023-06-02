<?php require('partials/head.php') ?>
<?php require('partials/nav.php') ?>
<?php require('partials/banner.php') ?>

    <main>
        <div class="mx-auto max-w-7xl py-6 sm:px-6 lg:px-8">
            <p class="mb-3">
                <a href="/notes/create" class="text-blue-500 hover:underline">Create Note</a>
            </p>

            <hr>

            <ul class="mt-3">
                <?php foreach ($notes as $note) : ?>
                    <li>
                        <a href="/note?id=<?= $note['id'] ?>" class="text-blue-500 hover:underline">
                            <?= htmlspecialchars($note['title']) ?>
                        </a>
                    </li>
                <?php endforeach; ?>
            </ul>
        </div>
    </main>

<?php require('partials/footer.php') ?>



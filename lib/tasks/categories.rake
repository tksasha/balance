namespace :categories do
  task fix_income: :environment do
    ActiveRecord::Base.connection.execute %q(UPDATE categories SET income=1 WHERE income='t')

    ActiveRecord::Base.connection.execute %q(UPDATE categories SET income=0 WHERE income='f')
  end
end

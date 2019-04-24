namespace :categories do
  task fix_boolean: :environment do
    # income
    ActiveRecord::Base.connection.execute %q(UPDATE categories SET income=1 WHERE income='t')

    ActiveRecord::Base.connection.execute %q(UPDATE categories SET income=0 WHERE income='f')

    # visible
    ActiveRecord::Base.connection.execute %q(UPDATE categories SET visible=1 WHERE visible='t')

    ActiveRecord::Base.connection.execute %q(UPDATE categories SET visible=0 WHERE visible='f')
  end
end

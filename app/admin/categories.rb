# frozen_string_literal: true

ActiveAdmin.register Category do
  menu priority: 3, label: proc { I18n.t('active_admin.categories') }

  actions :all, except: %i[destroy new create]

  filter :currency, as: :select, collection: CURRENCIES
end

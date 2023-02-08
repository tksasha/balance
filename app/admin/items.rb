# frozen_string_literal: true

ActiveAdmin.register Item do
  menu priority: 2, label: proc { I18n.t('active_admin.items') }

  actions :all, except: %i[destroy new create]

  permit_params %i[date]

  includes :category

  CURRENCIES.each_key do |currency|
    default = currency == CURRENCIES.keys.first

    scope(currency.upcase, default:, show_count: false) do |scope|
      scope.public_send(currency)
    end
  end

  filter :category, as: :select, collection: []
end

# frozen_string_literal: true

ActiveAdmin.register Category do
  include HasCurrencyScopes

  menu priority: 3, label: proc { I18n.t('active_admin.categories') }

  actions :all, except: %i[destroy]

  permit_params :name, :currency, :supercategory, :visible, :income

  filter :name

  controller do
    private

    def collection
      @collection ||=
        if request.format.json?
          Category.ransack(params[:q]).result.page(params[:page])
        else
          super
        end
    end
  end

  config.remove_action_item :new

  action_item :new, only: :index do
    link_to t(:create, scope: 'active_admin.buttons.categories'), new_admin_category_path
  end
end

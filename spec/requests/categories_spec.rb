# frozen_string_literal: true

RSpec.describe 'Categories' do
  %w[uah usd eur].each do |currency|
    describe 'GET index.js' do
      before do
        create_list(:category, 2, currency:)

        get "/#{ currency }/categories", xhr: true
      end

      it_behaves_like 'index.js'
    end

    describe 'GET new.js' do
      before { get "/#{ currency }/categories/new", xhr: true }

      it_behaves_like 'new.js'
    end

    describe 'POST create.js' do
      context 'when params are invalid' do
        let(:params) { { category: { name: '' } } }

        before { post "/#{ currency }/categories", params:, xhr: true }

        it_behaves_like 'new.js'
      end

      context 'when params are valid' do
        let(:name) { Faker::Commerce.product_name }

        let(:params) { { category: { name: } } }

        let(:category) { Category.last }

        before { post "/#{ currency }/categories", params:, xhr: true }

        it_behaves_like 'create.js'

        it { expect(category.name).to eq name }

        it { expect(category.currency).to eq currency }
      end
    end

    describe 'GET edit.js' do
      let(:category) { create(:category, currency:) }

      before { get "/#{ currency }/categories/#{ category.id }/edit", xhr: true }

      it_behaves_like 'edit.js'
    end

    describe 'PATCH update.js' do
      let(:category) { create(:category, currency:) }

      context 'when params are valid' do
        let(:name) { Faker::Commerce.product_name }

        let(:params) { { category: { name: } } }

        before do
          patch "/#{ currency }/categories/#{ category.id }", params:, xhr: true

          category.reload
        end

        it_behaves_like 'update.js'

        it { expect(category.name).to eq(name) }

        it { expect(category.currency).to eq(currency) }
      end
    end

    describe 'DELETE destroy.js' do
      let(:category) { create(:category, currency:) }

      before do
        delete "/#{ currency }/categories/#{ category.id }", xhr: true

        category.reload
      end

      it_behaves_like 'destroy.js'

      it { expect { Category.find(category.id) }.to raise_error(ActiveRecord::RecordNotFound) }
    end
  end
end

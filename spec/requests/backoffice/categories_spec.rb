# frozen_string_literal: true

RSpec.describe 'Backoffice::Categories', type: :request do
  %w[uah usd eur].each do |currency|
    describe 'GET index.js' do
      before do
        create_list(:category, 2, currency:)

        get "/#{ currency }/backoffice/categories", xhr: true
      end

      it_behaves_like 'index.js'
    end

    describe 'GET new.js' do
      before { get "/#{ currency }/backoffice/categories/new", xhr: true }

      it_behaves_like 'new.js'
    end

    describe 'POST create.js' do
      before { post "/#{ currency }/backoffice/categories", params:, xhr: true }

      context 'with valid params' do
        let(:params) do
          {
            category: {
              name: 'Drinks',
              supercategory: 'two',
              income: true,
              visible: true
            }
          }
        end

        let(:category) { Category.last }

        it_behaves_like 'create.js'

        it { expect(category.supercategory).to eq 'two' }

        it { expect(category.name).to eq 'Drinks' }

        it { expect(category.income).to be_truthy }

        it { expect(category.visible).to be_truthy }
      end

      context 'with invalid params' do
        let(:params) { { category: { name: '' } } }

        it_behaves_like 'new.js'
      end
    end

    describe 'GET edit.js' do
      let(:category) { create(:category, currency:) }

      before { get "/#{ currency }/backoffice/categories/#{ category.id }/edit", xhr: true }

      it_behaves_like 'edit.js'
    end

    describe 'PATCH update.js' do
      let(:category) do
        create(
          :category,
          name: 'First',
          supercategory: 'one',
          income: false,
          visible: false
        )
      end

      before { patch "/#{ currency }/backoffice/categories/#{ category.id }", params:, xhr: true }

      context 'with valid params' do
        let(:params) do
          {
            category: {
              name: 'Second',
              supercategory: 'two',
              income: true,
              visible: true
            }
          }
        end

        before { category.reload }

        it_behaves_like 'update.js'

        it { expect(category.name).to eq 'Second' }

        it { expect(category.supercategory).to eq 'two' }

        it { expect(category.income).to be_truthy }

        it { expect(category.visible).to be_truthy }
      end

      context 'with invalid params' do
        let(:params) { { category: { name: '' } } }

        it_behaves_like 'edit.js'
      end
    end
  end
end

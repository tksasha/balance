# frozen_string_literal: true

RSpec.describe 'Admin/Categories' do
  describe 'PATH update' do
    let(:category) do
      create(
        :category,
        name: 'Category #1',
        currency: 'uah',
        supercategory: 'one',
        income: false,
        visible: false
      )
    end

    before { patch "/admin/categories/#{ category.id }", params: }

    context 'with valid params' do
      let(:params) do
        {
          category: {
            name: 'Category #2',
            currency: 'usd',
            supercategory: 'two',
            income: true,
            visible: true
          }
        }
      end

      before { category.reload }

      it { is_expected.to redirect_to "/admin/categories/#{ category.id }" }

      it { expect(category.name).to eq 'Category #2' }
      it { expect(category.currency).to eq 'usd' }
      it { expect(category.supercategory).to eq 'two' }
      it { expect(category.income).to be_truthy }
      it { expect(category.visible).to be_truthy }
    end
  end
end

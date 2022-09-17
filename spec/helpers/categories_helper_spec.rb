# frozen_string_literal: true

RSpec.describe CategoriesHelper, type: :helper do
  describe '#supercategories' do
    subject { helper.supercategories }

    it do
      I18n.with_locale(:en) do
        expect(subject).to eq 'First' => 'first', 'Second' => 'second', 'Third' => 'third'
      end
    end

    it do
      I18n.with_locale(:ua) do
        expect(subject).to eq 'Перша' => 'first', 'Друга' => 'second', 'Третя' => 'third'
      end
    end
  end
end

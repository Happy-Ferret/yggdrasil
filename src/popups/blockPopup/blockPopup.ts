import { App } from '../../app.js'
import { Subscriber, Dispatcher } from '../../dispatcher.js'
import { AppEvent, LineStyle, MouseMode, Values } from '../../enums/enums.js'
import { Model } from '../../models/model.js'
import { Block } from '../../models/block.js'
import { BlockView } from '../../views/blockView.js'
import { Popup } from '../popups.js'
import { IdPopup, IdInput, IdTextarea, IdRange, IdLineStyle } from '../../controls/controls.js'

export class BlockPopup extends Popup implements Subscriber {
  private block: Block;
  private ctrlLineStyle: IdLineStyle;
  private ctrlLinewidth: IdRange;

  constructor() {
    super('blockpopup', Handlebars.templates.blockPopup, { colors: Values.COLORS_STANDARD });

    Dispatcher.subscribe(this);

    new IdPopup('.js-color', this.elem);
    new IdPopup('.js-line', this.elem);
    new IdPopup('.js-position', this.elem);
    new IdPopup('.js-delete', this.elem).addEventListener('click', () => { this.deleteBlock(); });
    new IdPopup('.js-more', this.elem).addEventListener('click', () => { this.showMore(); });    

    let btns = this.elem.querySelectorAll('.js-color id-popup');
    for(var i = 0; i < btns.length; i++) {
      let popup = new IdPopup(btns[i] as HTMLElement);
      let color = Values.COLORS_STANDARD[i];
      popup.backgroundColor = color;
      popup.addEventListener('click', () => { this.setColor(color); });
    }

    this.ctrlLineStyle = new IdLineStyle('.js-linestyle', this.elem).addEventListener('change', () => { this.block.lineStyle = this.ctrlLineStyle.value; });
    this.ctrlLinewidth = new IdRange('.js-linewidth', this.elem).addEventListener('input', () => { this.block.lineWidth = this.ctrlLinewidth.value; });

    this.elem.querySelector('.js-front').addEventListener('click', () => { 
      this.block.bringToFront();
      Dispatcher.notify(AppEvent.Load, null);
    });

    this.elem.querySelector('.js-forward').addEventListener('click', () => { 
      this.block.bringForward();
      Dispatcher.notify(AppEvent.Load, null);
    });

    this.elem.querySelector('.js-backward').addEventListener('click', () => { 
      this.block.sendBackward();
      Dispatcher.notify(AppEvent.Load, null);
    });

    this.elem.querySelector('.js-back').addEventListener('click', () => { 
      this.block.sendToBack();
      Dispatcher.notify(AppEvent.Load, null);
    });    
  }

  notify(event: AppEvent, model: Model) {
    if(event == AppEvent.MouseMove || event == AppEvent.Select) this.toggle();
  }  

  setColor(color: string) {
    this.block.fillColor = color;
  }

  deleteBlock() {
    App.pushUndo();
    this.block.delete();
    this.toggle();
  }

  showMore() {
    Dispatcher.notify(AppEvent.More, this.block);
  }

  toggle() {
    if(App.selection.isSingle() && App.selection.first() instanceof BlockView && App.mouseMode == MouseMode.None) {
      this.block = (App.selection.first().getModel() as Block);
      this.elem.style.left = App.canvas.offsetWidth / 2 + App.centerX + this.block.x * App.zoom + "px";
      this.elem.style.top = App.canvas.offsetHeight / 2 + App.centerY + this.block.y - 64 + "px";
      this.elem.style.display = 'flex';
      // Close any open overlays inside popup.
      let overlays = this.elem.querySelectorAll(".popup-overlay");
      for(let i = 0; i < overlays.length; i++) {
        (overlays[i] as HTMLElement).style.display = 'none';
      }
      this.ctrlLinewidth.value = this.block.lineWidth;
    } else {
      this.elem.style.display = 'none';
    }
  }  
}
